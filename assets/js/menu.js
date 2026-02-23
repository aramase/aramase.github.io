const menuTrigger = document.querySelector('.menu-trigger');
const menu = document.querySelector('.menu__inner--mobile');
const menuMoreTrigger = document.querySelector('.menu__sub-inner-more-trigger');
const menuMore = document.querySelector('.menu__sub-inner-more');

if (menuTrigger && menu) {
  menuTrigger.addEventListener('click', () => {
    menu.classList.toggle('hidden');
    const expanded = !menu.classList.contains('hidden');
    menuTrigger.setAttribute('aria-expanded', String(expanded));
  });
}

if (menuMoreTrigger && menuMore) {
  menuMoreTrigger.addEventListener('click', () => {
    menuMore.classList.toggle('hidden');
  });
}
