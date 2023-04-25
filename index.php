<?php

$api_php = true;

$parts = explode("/", $_SERVER["REQUEST_URI"]);

// gestion des redirections du Forum
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

// gestion des redirections de l'API en php si activé
if ($api_php) {
    if ($parts[1] == "apiForum") {
        include("api-php/api_php.php");
    }
}