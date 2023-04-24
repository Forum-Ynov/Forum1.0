<?php

$api_php = false;

$parts = explode("/", $_SERVER["REQUEST_URI"]);

if ($parts[1] == "Forum") {
    if ($parts[2] == "home") {
        include("site/Static/Html/home.php");
    } elseif ($parts[2] == "log_in") {
        include("site/Static/Html/log_in.php");
    } elseif ($parts[2] == "sign_up") {
        include("site/Static/Html/sign_up.php");
    } elseif ($parts[2] == "profil") {
        include("site/Static/Html/profil.php");
    } else {
        header("Location: /Forum/home");
    }
}

if ($api_php) {
    if ($parts[1] == "apiForum") {
        include("api-php/api_php.php");
    }
}
