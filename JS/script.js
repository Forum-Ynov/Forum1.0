
import { User } from "/JS/user_class.js"

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
            console.log(storageUser)

            const r = await fetch(`http://localhost:8000/apiForum/users/${storageUser.id_user}`, {
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

            const r = await fetch(`http://localhost:8000/apiForum/users/${storageUser.id_user}`, {
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

