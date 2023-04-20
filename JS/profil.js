import { User } from "./user_class.js"

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

} else {
    console.log("to connect")
}
