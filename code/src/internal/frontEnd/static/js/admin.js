// Ouverture des modales
document.getElementById('btn-moderation').addEventListener('click', function() {
    document.getElementById('modal-moderation').classList.remove('hidden');
    initializePagination('tab-content-waiting', '.waiting-content-item', '.pagination-info-waiting', '.pagination-waiting', 4);
    initializePagination('tab-content-moderation', '.moderation-report-item', '.pagination-info-moderation', '.pagination-moderation', 2);
});

document.getElementById('btn-reports').addEventListener('click', function() {
    document.getElementById('modal-reports').classList.remove('hidden');
    initializePagination('user-roles-tab', '.user-role-item', '.pagination-info-roles', '.pagination-controls-roles', 2);
});

document.getElementById('btn-category').addEventListener('click', function() {
    document.getElementById('modal-category').classList.remove('hidden');
    initializePagination('modal-category', '.category-item', '.pagination-info', '.pagination-controls', 6);
});

document.getElementById('btn-add-category').addEventListener('click', function() {
    document.getElementById('modal-add-category').classList.remove('hidden');
});

// Fermeture des modales
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
document.querySelectorAll('.tab-btn-moderation, .tab-btn-waiting').forEach(button => {
    button.addEventListener('click', () => {
        const tabId = button.getAttribute('data-tab');
        document.querySelectorAll('.tab-btn-moderation, .tab-btn-waiting').forEach(btn => {
            btn.classList.remove('active-tab');
            btn.classList.add('bg-gray-600');
        });
        button.classList.add('active-tab');
        button.classList.remove('bg-gray-600');
        document.querySelectorAll('.tab-content-moderation, .tab-content-waiting').forEach(content => {
            content.classList.add('hidden');
        });
        document.getElementById(tabId).classList.remove('hidden');
        // Réinitialise la pagination à chaque changement d’onglet
        if (tabId === 'tab-content-moderation') {
            initializePagination('tab-content-moderation', '.moderation-report-item', '.pagination-info-moderation', '.pagination-moderation', 2);
        } else if (tabId === 'tab-content-waiting') {
            initializePagination('tab-content-waiting', '.waiting-content-item', '.pagination-info-waiting', '.pagination-waiting', 4);
        }
    });
});

document.addEventListener('DOMContentLoaded', function() {
    // Rien ici, tout est géré à l'ouverture des modales
});

function openEditCategoryPopup(categoryId) {
    const modal = document.getElementById('edit-category-popup-' + categoryId);
    if (modal) {
        modal.classList.remove('hidden');
    }
}