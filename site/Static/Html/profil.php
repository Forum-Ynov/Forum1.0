<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/site/Static/CSS/profil.css">
    <title>Profil</title>
</head>

<body data-theme="light">
    <header>
        <nav class="navbar">
            <a href="/Forum/home">
                <img src="/site/Assets/Images/others/logo_forum_light.png" alt="logo" class="logo">
                <img src="/site/Assets/Images/others/logo_forum.png" alt="logo_dark" class="logo" style="opacity: .3;">
            </a>
            <label class="switch">
                <input type="checkbox" name="switch_theme" id="swmode">
                <span class="slider"></span>
            </label>
            <button onclick="toggleMenu()" class="burger"></button>
            <div class="dropdowns">
                <div class="dropdown">
                    <div id="profil_pseudo"></div>
                    <button class="button_user_img">
                        <img src="/site/Assets/Images/profil/user.svg" alt="user">
                        <img src="/site/Assets/Images/menu/arrow_down.svg" />
                    </button>
                    <div class="dropdown-menu">
                        <a href="/Forum/profil"><button class="button_user">profil</button></a>
                        <button class="button_user">paramètres</button>
                        <button class="button_user" id="btn_deconenxion">déconnexion</button>
                    </div>
                </div>
            </div>
        </nav>
    </header>
    <div class="containers">
        <div class="top_info">
            <div class="left_top_info">
                <img id="user_pp" src="/site/Assets/Images/profil/chat.svg" alt="profil">
                <h3 id="user_pseudo">user121519</h3>
            </div>
        </div>
        <div class="info">
            <div class="left_col_info">
                <div class="name_info">
                    <h5>Nom d'utilisateur :</h5>
                    <input value="user121519" disabled id="name"></input>
                </div>
                <div class="mail_info">
                    <h5>Mail :</h5>
                    <input value="User@ynov.com" disabled id="mail"></input>
                </div>
                <div class="mdp_info">
                    <h5>Mot de passe :</h5>
                    <h3><input type="password" disabled value="CtropBienLesInputs" id="password"></input></h3>
                </div>
            </div>
            <div class="right_col_info">
                <div class="btn_name_info"><button class="button_simple_update" id="update_name">Modifier le
                        nom</button></div>
                <div class="btn_mail_info"><button class="button_simple_update" id="update_mail">Modifier le
                        mail</button></div>
                <div class="btn_mdp_info"><button class="button_simple_update" id="update_mdp">Modifier le mdp</button>
                </div>
            </div>
        </div>
    </div>
</body>

<script src="https://code.jquery.com/jquery-3.6.4.min.js"></script>
<script type="module" src="/site/JS/script.js"></script>
<script type="module" src="/site/JS/profil.js"></script>

</html>