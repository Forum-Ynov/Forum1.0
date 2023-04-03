const toggleMenu = () => document.body.classList.toggle("open");

const body = document.querySelector('body');
const logo_dark = document.getElementById('logo_dark')
const logo_light = document.getElementById('logo_light')

console.log(logo_dark);
console.log(logo_light);

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


