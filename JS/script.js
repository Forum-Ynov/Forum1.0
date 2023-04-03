const toggleMenu = () => document.body.classList.toggle("open");

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
        logo_light.style.opacity = 0;
        logo_dark.style.opacity = 1;
    }
    else {
        console.log("light");
        body.setAttribute('data-theme', 'light');
        logo_light.style.opacity = 1;
        logo_dark.style.opacity = 0;

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
