const toggleMenu = () => document.body.classList.toggle("open");

const body = document.querySelector('body');
function switch_theme() {
    if (test.checked) {
        body.setAttribute('data-theme', 'dark');
    }
    else {
        body.setAttribute('data-theme', 'light');
    }
}