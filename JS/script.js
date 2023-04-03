
import { User } from "./user_class.js"

const toggleMenu = () => document.body.classList.toggle("open");

const swmode = document.getElementById("swmode")

const body = document.querySelector('body');
const logo_dark = document.getElementById('logo_dark')
const logo_light = document.getElementById('logo_light')
let openpop = document.getElementsByName('openpop');
let Popup = document.getElementById('favDialog');
let close_popup = document.getElementById('close_pop')
let content = document.getElementById('contenu')

function switch_theme() {
    if (test.checked) {
        console.log("dark");
        body.setAttribute('data-theme', 'dark');

        const localUser = localStorage.getItem("loged_user")?.toString()
        let storageUser = new User("", "", "", "", "", "");

        if (localUser) {
            storageUser = JSON.parse(localUser)
            console.log(storageUser)

            const r = await fetch(`http://localhost:8000/usertheme/${storageUser.id_user}`, {
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

            const r = await fetch(`http://localhost:8000/usertheme/${storageUser.id_user}`, {
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

function openpopup() {
    Popup.showModal();
    content.style.position = 'fixed';
};
close_popup.addEventListener('click', function onClose() {
    Popup.close();
    content.style.position = 'initial'
});
