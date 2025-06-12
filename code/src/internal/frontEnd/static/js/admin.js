// Open modal management

document.getElementById('btn-moderation').addEventListener('click', function() {
    document.getElementById('modal-moderation').classList.remove('hidden');
});

document.getElementById('btn-reports').addEventListener('click', function() {
    document.getElementById('modal-reports').classList.remove('hidden');
});

document.getElementById('btn-category').addEventListener('click', function() {
    document.getElementById('modal-category').classList.remove('hidden');
});

document.getElementById('btn-add-category').addEventListener('click', function() {
    document.getElementById('modal-add-category').classList.remove('hidden');
}
);

// Close modal management
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

// Tab management
document.querySelectorAll('.tab-btn-moderation, .tab-btn-waiting').forEach(button => {
    button.addEventListener('click', () => {
        const tabId = button.getAttribute('data-tab');

        // Active tab button
        document.querySelectorAll('.tab-btn-moderation, .tab-btn-waiting').forEach(btn => {
            btn.classList.remove('active-tab');
            btn.classList.add('bg-gray-600');
        });
        button.classList.add('active-tab');
        button.classList.remove('bg-gray-600');

        // Show corresponding content
        document.querySelectorAll('.tab-content-moderation, .tab-content-waiting').forEach(content => {
            content.classList.add('hidden');
        });
        document.getElementById(tabId).classList.remove('hidden');
    });
});

document.addEventListener('DOMContentLoaded', function() {
    // Pagination for the moderation modal ("Pending" tab)
    document.getElementById('btn-moderation').addEventListener('click', function() {
        document.getElementById('modal-moderation').classList.remove('hidden');
        // Pagination on pending items
        initializePagination(
            'tab-content-waiting',
            '.waiting-content-item',
            '.pagination-info-waiting',
            '.pagination-waiting',
            4
        );
        // Pagination on reports
        initializePagination(
            'tab-content-moderation',
            '.moderation-report-item',
            '.pagination-info-moderation',
            '.pagination-moderation',
            2
        );

    initializePagination('reported-users-tab', '.user-report-item', '.pagination-info-users', '.pagination-controls-users', 2);
    initializePagination('modal-category', 'li', '.pagination-info', '.pagination-controls', 6);
    document.getElementById('btn-moderation').addEventListener('click', function() {
        document.getElementById('modal-moderation').classList.remove('hidden');
    });

    // Pagination for user management
    document.getElementById('btn-reports').addEventListener('click', function() {
        document.getElementById('modal-reports').classList.remove('hidden');
    });

    // Pagination for categories management
    document.getElementById('btn-category').addEventListener('click', function() {
        document.getElementById('modal-category').classList.remove('hidden');
    });

    document.querySelector('[data-tab="user-roles-tab"]').addEventListener('click', function() {
        initializePagination('user-roles-tab', '.user-role-item', '.pagination-info-roles', '.pagination-controls-roles', 2);
    });
});
})

function openEditCategoryPopup(categoryId) {
    const modal = document.getElementById('edit-category-popup-' + categoryId);
    if (modal) {
        modal.classList.remove('hidden');
    }
}
