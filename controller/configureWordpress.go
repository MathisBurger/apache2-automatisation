package controller

import (
	"database/sql"
	"fmt"
	"github.com/MathisBurger/apache2-automatisation/utils"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type configureWordpressResponse struct {
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
	Status     string `json:"status"`
	Error      string `json:"error"`
}

type pwdStruct struct {
	Password string `json:"password"`
}

// controller to configure wordpress installation on apache2
func ConfigureWordpressController(c *fiber.Ctx) error {

	// check custom cors
	if !utils.CheckCORS(c.IP()) {
		return c.JSON(configureWordpressResponse{
			"Your origin is not allowed",
			200,
			"ok",
			"None",
		})
	}

	domain := c.Query("domain")
	AuftragsID := c.Query("AuftragsID")

	// define paths
	docPath := "/var/www/" + domain
	cfgPath := "/etc/apache2/sites-available/" + domain + ".conf"

	// create config path if not exists
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		_, err := os.Create(cfgPath)
		if err != nil {
			return c.JSON(configureWordpressResponse{
				"Error while installing wordpress",
				200,
				"ok",
				err.Error(),
			})
		}
	}

	// create document root if not exists
	if _, err := os.Stat(docPath); os.IsNotExist(err) {
		err := os.Mkdir(docPath, 0755)
		if err != nil {
			return c.JSON(configureWordpressResponse{
				"Error while installing wordpress",
				200,
				"ok",
				err.Error(),
			})
		}
	}

	conn := utils.GetConn()

	// counting number of databases
	stmt, err := conn.Prepare("SELECT `ID` FROM `active_databases` WHERE `Name` LIKE ?;")
	if err != nil {
		fmt.Println("Error index: 0")
		panic(err.Error())
	}
	resp, err := stmt.Query("aaa_" + AuftragsID + "_%")
	if err != nil {
		fmt.Println("Error index: 1")
		panic(err.Error())
	}
	counter := 0
	for resp.Next() {
		counter += 1
	}

	// getting database name
	DatabaseName := NameCalculator(counter, AuftragsID)

	// At this position allowed (will be fixed later)
	stmt, err = conn.Prepare("CREATE DATABASE " + DatabaseName + ";")
	if err != nil {
		return c.JSON(configureWordpressResponse{
			"Error while installing wordpress",
			200,
			"ok",
			err.Error(),
		})
	}
	_, err = stmt.Exec()
	if err != nil {
		return c.JSON(configureWordpressResponse{
			"Error while installing wordpress",
			200,
			"ok",
			err.Error(),
		})
	}

	// insert new database to database list
	stmt, err = conn.Prepare("INSERT INTO `active_databases` (`ID`, `Name`, `last-edited`) VALUES (NULL, ?, CURRENT_TIMESTAMP());")
	if err != nil {
		return c.JSON(configureWordpressResponse{
			"Error while installing wordpress",
			200,
			"ok",
			err.Error(),
		})
	}

	_, err = stmt.Exec(DatabaseName)

	if err != nil {
		return c.JSON(configureWordpressResponse{
			"Error while installing wordpress",
			200,
			"ok",
			err.Error(),
		})
	}

	// copies wordpress installation to document root
	utils.Copy("/var/www/software/wordpress", docPath)

	// writes config
	data, _ := ioutil.ReadFile("/root/automatisation/InstallationService/sample/wordpress.conf")
	modified := []byte(strings.ReplaceAll(string(data), "{{DOMAIN}}", domain))
	err = ioutil.WriteFile(cfgPath, modified, 0644)

	if err != nil {
		return c.JSON(configureWordpressResponse{
			"Error while installing wordpress",
			200,
			"ok",
			err.Error(),
		})
	}
	stmt.Close()

	// configures config.php in installation (mysql credentials)
	Configure_WR_Config(conn, DatabaseName, docPath)
	conn.Close()
	return c.JSON(configureWordpressResponse{
		"Successfully installed wordpress",
		200,
		"ok",
		"None",
	})
}

// calculates database name
func NameCalculator(counter int, AuftragsID string) string {
	index := counter + 1
	resp := "aaa_" + AuftragsID + "_"
	if index < 10 {
		return resp + "0" + strconv.Itoa(index)
	} else {
		return resp + strconv.Itoa(index)
	}
}

// configures wr-config.php for installation
func Configure_WR_Config(conn *sql.DB, DatabaseName, docPath string) {
	data, _ := ioutil.ReadFile("/root/automatisation/InstallationService/sample/wordpress.php")
	dbuser := "aaa_" + strings.Split(DatabaseName, "_")[1]
	stmt, _ := conn.Prepare("SELECT `password` FROM `database_accounts` WHERE `username`=?")
	resp, err := stmt.Query(dbuser)
	if err != nil {
		panic(err)
	}
	var pwd pwdStruct
	for resp.Next() {
		err = resp.Scan(&pwd.Password)
		if err != nil {
			panic(err)
		}
	}
	modified := []byte(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(data),
		"{{DatabaseName}}", DatabaseName),
		"{{DatabaseUser}}", dbuser),
		"{{DatabasePassword}}", pwd.Password))
	cfgPath := docPath + "/wp-config.php"
	os.Create(cfgPath)
	ioutil.WriteFile(cfgPath, modified, 0644)
	resp.Close()
	stmt.Close()
}
