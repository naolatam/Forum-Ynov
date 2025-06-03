// Gestion du menu mobile
document.addEventListener('DOMContentLoaded', function () {
    const mobileMenuButton = document.getElementById('mobile-menu-button');
    const mobileMenu = document.getElementById('mobile-menu');

    mobileMenuButton.addEventListener('click', function () {
        mobileMenu.classList.toggle('hidden');
    });

    // Fermer le menu quand on clique ailleurs
    document.addEventListener('click', function (event) {
        if (!mobileMenuButton.contains(event.target) &&
            !mobileMenu.contains(event.target) &&
            !mobileMenu.classList.contains('hidden')) {
            mobileMenu.classList.add('hidden');
        }
    });
});
