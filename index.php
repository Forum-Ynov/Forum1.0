<?php

$parts = explode("/", $_SERVER["REQUEST_URI"]);

if ($parts[1] == "home") {
    include("Static/Html/home.php");
} elseif ($parts[1] == "log_in") {
    include("Static/Html/log_in.php");
} elseif ($parts[1] == "sign_up") {
    include("Static/Html/sign_up.php");
} elseif ($parts[1] == "profil") {
    include("Static/Html/profil.php");
} else {
    header("Location: /home");
}