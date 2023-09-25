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
    <link rel="stylesheet" href="/site/Static/CSS/home.css">
    <title>Home page</title>
</head>

<style id="style_mod">
    .tagsall::before {
        content: url(/site/Assets/Images/icon_input/arobase.svg);
        position: absolute;
        left: -20px;
    }
</style>

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
            <button class="button" id="open_create">Log in</button>
            <div class="dropdowns">
                <div class="dropdown">
                    <div id="profil_pseudo"></div>
                    <button class="button_user_img">
                        <img id="profil_picture" src="/site/Assets/Images/profil/user.svg" alt="user">
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
    <aside class="sidebar">
        <div class="div_tags">
            <h3>Tags:</h3>
            <div class="element_tag" id="display_tags">

            </div>
        </div>
        <div class="div_text">
            <p class="info_aside">Connectez vous pour profiter des fonctionnalités de création de post, réponse au post
                ... et bien plus encore</p>
        </div>
        <div class="div_text_connect">
            <p class="info_aside">Pour créer un topic, cliquez sur le bouton "Create Post" et remplissez les
                informations demandées. C'est simple !</p>
        </div>
        <div class="div_button_aside">
            <hr>
            <button class="button" id="myBtn">Create Post</button>
        </div>
    </aside>

    <article class="content" id="display_topics">

    </article>
    <dialog id="favDialog" class="popup">
        <div class="container">
            <div class="box">
                <div class="part_left">
                    <div id="left_container">
                    </div>
                    <form action="" id="CreateMessage" class="input_comment">
                        <textarea name="comment" id="comment" cols="50" rows="10"
                            placeholder="Write a comment..."></textarea>
                        <button type="submit" class="button_pop">Publier</button>
                    </form>
                </div>
                <div class="part_right">
                    <img src="/site/Assets/Images/menu/close.svg" alt="close" id="close_pop">
                    <div class="comment">
                        <div class="comment_top">
                            <p>Commentaire :</p>
                        </div>
                        <div class="comment_main" id="comment_container">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </dialog>
    <form action="" id="Createpost" class="popup_create">
        <div class="container_create">
            <div class="post_close">
                <img src="/site/Assets/Images/menu/close.svg" alt="fermeture popup" />
            </div>
            <h3>Create a new post</h3>
            <label for="" class="label_title">Titre :</label>
            <textarea name="title" id="post-title" cols="30" rows="10" class="area_title"></textarea>
            <label for="" class="label_title">Description :</label>
            <textarea name="description" id="post-description" cols="30" rows="10" class="area_description"></textarea>
            <div class="input_tag">
                <select id="tag-dropdown">
                    <option value="">Select a tag</option>
                </select>
            </div>
            <div id="status-message"></div>
            <button type="submit" class="button_create">Publier</button>
        </div>
    </form>

</body>

<script src="https://code.jquery.com/jquery-3.6.4.min.js"></script>
<script type="module" src="/site/JS/script.js"></script>
<script type="module" src="/site/JS/home.js"></script>
<script type="module">
    import { openmessage } from "/site/JS/home.js"
    window.openmessage = openmessage
</script>

</html>