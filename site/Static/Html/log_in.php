<!DOCTYPE html>
<html lang="fr">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link
        href="https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;1,100&display=swap"
        rel="stylesheet">
    <link rel="stylesheet" href="/site/Static/CSS/log.css">
    <title>Log in</title>
</head>

<body>
    <nav class="navbar">
        <a href="/Forum/home">
            <img src="/site/Assets/Images/others/logo_forum_light.png" alt="logo" class="logo">
            <img src="/site/Assets/Images/others/logo_forum.png" alt="logo_dark" class="logo" style="opacity: .3;">
        </a>
        <label class="switch">
            <input type="checkbox" name="switch_theme" id="swmode">
            <span class="slider"></span>
        </label>
    </nav>
    <div class="container">
        <h2>Log in to your account</h2>
        <form action="" class="formulaire" id="form_log">
            <input id="in_pseudo" type="text" placeholder="Nom utilisateur" required name="pseudo">
            <div class="div_input"><input id="in_passwd" type="password" placeholder="Mot de passe" required
                    name="passwd">
                <span class="img_input" id="passtotext">
                    <img src="/site/Assets/Images/icon_input/eye.svg" alt="eye" id="imgpass">
                </span>
            </div>

            <button type="submit">Log In</button>

        </form>
        <div id="message">

        </div>
        <a href="/Forum/sign_up">
            <p>Vous n'avez pas de compte inscrivez vous !!!</p>
        </a>
    </div>

</body>

<script src="https://code.jquery.com/jquery-3.6.4.min.js"></script>
<script type="module" src="/site/JS/log_in.js"></script>
<script type="module" src="/site/JS/script.js"></script>

</html>