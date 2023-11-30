
import { User } from "/site/JS/user_class.js"

let api_port

$.getJSON("/site/conf.json", function (data) {
    api_port = data.api_port
})

const toggleMenu = () => document.body.classList.toggle("open");

const swmode = document.getElementById("swmode")
const body = document.querySelector('body');
// const heart = document.getElementById('like');
// const heart_red = document.getElementById('like_red')
// console.log(heart);

swmode.onclick = async function switch_theme() {
    if (swmode.checked) {
        console.log("dark");
        body.setAttribute('data-theme', 'dark');

        const localUser = localStorage.getItem("loged_user")?.toString()
        let storageUser = new User("", "", "", "", "", "");

        if (localUser) {
            storageUser = JSON.parse(localUser)

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

