import { User } from "/site/JS/user_class.js";
import { Imagepp } from "/site/JS/pp_class.js";

let api_port

const profil_pseudo = document.getElementById("profil_pseudo");
const user_pp = document.getElementById("user_pp");
const user_pseudo = document.getElementById("user_pseudo");
const swmode = document.getElementById("swmode")
const localUser = localStorage.getItem("loged_user")?.toString()
let storageUser = new User("", "", "", "", "", "");


if (localUser) {
    console.log("auto connect")
    storageUser = JSON.parse(localUser)
    console.log(storageUser)

    switch (storageUser.theme) {
        case ("dark"):
            document.querySelector('body').setAttribute('data-theme', 'dark');
            console.log("default dark")
            swmode.checked = true
            break
        case ("light"):
            document.querySelector('body').setAttribute('data-theme', 'light');
            console.log("default light")
            swmode.checked = false
            break
    }

    $.getJSON("/site/conf.json", function (data) {
        api_port = data.api_port
        log_In()
    })

} else {
    document.location.href = "/Forum/home";
    console.log("to connect")
}

const btn_deconenxion = document.getElementById("btn_deconenxion")

btn_deconenxion.addEventListener("click", () => {
    localStorage.removeItem("loged_user")
    document.location.href = "/Forum/home";
});



function log_In() {
    if (localUser) {
        profil_pseudo.innerHTML = storageUser.pseudo;
        user_pseudo.innerHTML = storageUser.pseudo;
        user_pp.src = "/site/Assets/Images/profil/homer.svg"

        const ppload = fetch(window.location.origin+`/apiForum/pp/${storageUser.id_imagepp}`, {
            method: 'GET',
            headers: {
                "Accept": "application/json",
                "Content-type": "application/json; charset=UTF-8"
            }
        })
            .then((res) => {
                if (res.ok) {
                    res.json().then(data => {
                        user_pp.src = `/site/Assets/Images/profil/${data.image_loc}`
                    })
                } else {
                    console.log("res.ok false")
                }
            });



    }
}



















const update_name = document.getElementById('update_name')
const update_mail = document.getElementById('update_mail')
const update_mdp = document.getElementById('update_mdp')

const name_user = document.getElementById('name')
const mail = document.getElementById('mail')
const password = document.getElementById('password')



let clics_name = 0
update_name.addEventListener('click', event => {
    name_user.disabled = !name_user.disabled; // Alterne entre enabled et disabled à chaque clic
    clics_name++;
    if (clics_name % 2 !== 0) {
        update_name.innerHTML = "Valider"
        update_name.style.backgroundColor = "var(--couleur_secondaire)"
        update_name.style.color = "white"
        name_user.disabled = false;

    } else {
        update_name.innerHTML = "Modifier le nom"
        update_name.style.backgroundColor = "var(--couleur_principale)"
        update_name.style.color = "var(--couleur_secondaire)"
        name_user.disabled = true;


    }
});



let clics_mail = 0
update_mail.addEventListener('click', event => {
    mail.disabled = !mail.disabled; // Alterne entre enabled et disabled à chaque clic
    clics_mail++;
    if (clics_mail % 2 !== 0) {
        update_mail.innerHTML = "Valider"
        update_mail.style.backgroundColor = "var(--couleur_secondaire)"
        update_mail.style.color = "white"
        password.disabled = false;

    } else {
        update_mail.innerHTML = "Modifier le mail"
        update_mail.style.backgroundColor = "var(--couleur_principale)"
        update_mail.style.color = "var(--couleur_secondaire)"
        mail.disabled = true;
    }
});

let clics_passeword = 0
update_mdp.addEventListener('click', event => {
    password.disabled = !password.disabled; // Alterne entre enabled et disabled à chaque clic
    clics_passeword++;
    if (clics_passeword % 2 !== 0) {
        update_mdp.innerHTML = "Valider"
        update_mdp.style.backgroundColor = "var(--couleur_secondaire)"
        update_mdp.style.color = "white"
        password.disabled = false;
        password.type = "text"
    } else {
        update_mdp.innerHTML = "Modifier le mdp"
        update_mdp.style.backgroundColor = "var(--couleur_principale)"
        update_mdp.style.color = "var(--couleur_secondaire)"
        password.disabled = true;
        password.type = "password"
    }
});




const body = document.querySelector('body');


swmode.onclick = async function switch_theme() {
    if (swmode.checked) {
        console.log("dark");
        body.setAttribute('data-theme', 'dark');

        const localUser = localStorage.getItem("loged_user")?.toString()
        let storageUser = new User("", "", "", "", "", "");

        if (localUser) {
            storageUser = JSON.parse(localUser)
            console.log(storageUser)

            const r = await fetch(window.location.origin+`/apiForum/users/${storageUser.id_user}`, {
                method: 'PATCH',
                headers: {
                    "Accept": "application/json",
                    "Content-type": "application/json; charset=UTF-8"
                },
                body: JSON.stringify({ theme: "dark" })
            })
            storageUser.theme = "dark"
            localStorage.setItem("loged_user", JSON.stringify(storageUser))

        }

    }
    else {
        console.log("light");
        body.setAttribute('data-theme', 'light');

        const localUser = localStorage.getItem("loged_user")?.toString()
        let storageUser = new User("", "", "", "", "", "");

        if (localUser) {
            storageUser = JSON.parse(localUser)
            console.log(storageUser)

            const r = await fetch(window.location.origin+`/apiForum/users/${storageUser.id_user}`, {
                method: 'PATCH',
                headers: {
                    "Accept": "application/json",
                    "Content-type": "application/json; charset=UTF-8"
                },
                body: JSON.stringify({ theme: "light" })
            })
            storageUser.theme = "light"
            localStorage.setItem("loged_user", JSON.stringify(storageUser))

        }

    }
}

