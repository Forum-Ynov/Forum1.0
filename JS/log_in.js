
import { User } from "/JS/user_class.js"

let message = document.getElementById("message");
let in_pseudo = document.getElementById("in_pseudo");
let in_passwd = document.getElementById("in_passwd");
let form_log = document.getElementById("form_log");

let actual_user = new User("", "", "", "", "", "");
let default_user = new User("", "", "", "", "", "");

form_log.addEventListener("submit", async function (event) {
    event.preventDefault();
    let names = in_pseudo.value;
    let pass = in_passwd.value;

    const r = await fetch("https://api.sqwado.websr.fr/apiForum/login", {
        method: 'POST',
        headers: {
            "Accept": "application/json",
            "Content-type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({ pseudo: names, passwd: pass })
    })
        .then((res) => {
            console.log(res)
            if (res.ok) {
                console.log("res.ok true")
                res.json().then(data => {
                    actual_user.id_user = data.id_user;
                    actual_user.pseudo = data.pseudo;
                    actual_user.email = data.email;
                    actual_user.passwd = data.passwd;
                    actual_user.id_imagepp = data.id_imagepp;
                    actual_user.theme = data.theme
                    console.log(actual_user)
                    localStorage.setItem("loged_user", JSON.stringify(actual_user))
                    window.location.href = "/home"
                })
            } else {
                console.log("res.ok false")
                res.json().then(data => {
                    console.log("ERREUR");
                    message.innerHTML = data.message;
                    actual_user.id_user = ""
                    actual_user.pseudo = ""
                    actual_user.email = ""
                    actual_user.passwd = ""
                    actual_user.id_imagepp = ""
                    actual_user.theme = ""
                })
            }
        });
});

const swmode = document.getElementById("swmode")
const localUser = localStorage.getItem("loged_user")?.toString()

if (localUser) {
    document.location.href = "/home";
} else {
    console.log("to connect")
}


let passtotext = document.getElementById("passtotext")
let imgpass = document.getElementById("imgpass")

passtotext.onclick = function () {
    if (in_passwd.type == "password") {
        in_passwd.type = "text"
        imgpass.src = "/Assets/Images/icon_input/eye-off.svg"
    } else {
        in_passwd.type = "password"
        imgpass.src = "/Assets/Images/icon_input/eye.svg"
    }

}