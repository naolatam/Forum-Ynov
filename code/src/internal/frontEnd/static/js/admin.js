// Gestion des modals

document.getElementById('btn-moderation').addEventListener('click', function() {
    document.getElementById('modal-moderation').classList.remove('hidden');
});

document.getElementById('btn-signalements').addEventListener('click', function() {
    document.getElementById('modal-signalements').classList.remove('hidden');
});

// Fermeture des modals
document.querySelectorAll('.modal-close').forEach(function(button) {
    button.addEventListener('click', function() {
        document.getElementById(this.dataset.modal).classList.add('hidden');
    });
});

document.querySelectorAll('.absolute[id$="-backdrop"]').forEach(function(backdrop) {
    backdrop.addEventListener('click', function() {
        this.parentElement.classList.add('hidden');
    });
});

// Gestion des tabs
document.querySelectorAll('.tab-button').forEach(function(button) {
    button.addEventListener('click', function() {
        // Désactiver tous les onglets
        document.querySelectorAll('.tab-button').forEach(function(btn) {
            btn.classList.remove('active-tab', 'bg-gray-700');
            btn.classList.add('bg-gray-600');
        });

        // Activer l'onglet cliqué
        this.classList.add('active-tab', 'bg-gray-700');
        this.classList.remove('bg-gray-600');

        // Masquer tous les contenus
        document.querySelectorAll('.tab-content').forEach(function(content) {
            content.classList.add('hidden');
        });

        // Afficher le contenu correspondant
        document.getElementById(this.dataset.tab).classList.remove('hidden');
    });
});