<?php
define('FORCE_SSL_ADMIN', true);
 if ($_SERVER['HTTP_X_FORWARDED_PROTO'] == 'https')
        $_SERVER['HTTPS']='on';

define( 'DB_NAME', '{{DatabaseName}}' );
define( 'DB_USER', '{{DatabaseUser}}' );
define( 'DB_PASSWORD', '{{DatabasePassword}}' );
define( 'DB_HOST', '10.11.0.3' );
define( 'DB_CHARSET', 'utf8mb4' );
define('DB_COLLATE', '');
define('WP_MEMORY_LIMIT', '50M');
define( 'AUTH_KEY',         '&cBm_/i4LQq,iLo*7jC5y|;sK)kZ[(&PW/Fp.Ib 1QHhQUQ+zR46_yvsR&U.dbGf' );
define( 'SECURE_AUTH_KEY',  'GO[q9OJ0cw2S@#I;@44J8X~P~<C!b&AVo|<Q|{31J,MpF^&YZ+<PDP7+*|_pT1E2' );
define( 'LOGGED_IN_KEY',    '{B#h2oWo%tWOk.stfT=mx*~v E b5JWF ScWN=W}G`p`e_Oi&$E0@rZQ1R[~Oa3o' );
define( 'NONCE_KEY',        '::%T:m37uv}Y|qML.0juoj=Mnh11}?54fPd%/3x9q;C5qCA:!, m,M%qEJ@ s|Fr' );
define( 'AUTH_SALT',        '|]>h5X%|/p6PF]Si{/bKD_6`0XZGe#r^=EP#%IFT;+7}z4D/1}G.7geHE8CfJPf@' );
define( 'SECURE_AUTH_SALT', 'F;%K>5Z: 1$LMhqCu}9gBjwe6s-w9b!oW?m%w%(PQ:h7I4e[;.e7Z__AiQ#>KUp[' );
define( 'LOGGED_IN_SALT',   '4v`qgXP2pN1E_GR+TeTvL6*UABwE8TuQZB**mONp]DpKCEG;%0ENT:i/]lB2Ne#&' );
define( 'NONCE_SALT',       ');,SwoZPDlQd?R-u^.CU5;9cVvbEAH+L,h?DdNO]~+I1KhPL^Y!!GGa0o<YWtLx3' );
$table_prefix = 'wp_';
define( 'WP_DEBUG', false );
if ( ! defined( 'ABSPATH' ) ) {
	define( 'ABSPATH', dirname( __FILE__ ) . '/' );
}
require_once( ABSPATH . 'wp-settings.php' );
