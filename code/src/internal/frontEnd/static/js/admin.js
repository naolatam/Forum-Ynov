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
document.querySelectorAll('.tab-button').forEach(function(button) {
    button.addEventListener('click', function() {
        // Change tab style to disable
        document.querySelectorAll('.tab-button').forEach(function(btn) {
            btn.classList.remove('active-tab', 'bg-gray-700');
            btn.classList.add('bg-gray-600');
        });

        // Active tab style
        this.classList.add('active-tab', 'bg-gray-700');
        this.classList.remove('bg-gray-600');

        // Hide all tab contents
        document.querySelectorAll('.tab-content').forEach(function(content) {
            content.classList.add('hidden');
        });

        // Show the selected tab content
        document.getElementById(this.dataset.tab).classList.remove('hidden');
    });
});

document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('btn-moderation').addEventListener('click', function() {
        document.getElementById('modal-moderation').classList.remove('hidden');
        initializePagination('content-tab', '.content-item', '.pagination-info-content', '.pagination-controls-content', 4);
    });

    document.getElementById('btn-reports').addEventListener('click', function() {
        document.getElementById('modal-reports').classList.remove('hidden');
        initializePagination('reported-users-tab', '.user-report-item', '.pagination-info-users', '.pagination-controls-users', 2);
    });

    document.getElementById('btn-category').addEventListener('click', function() {
        document.getElementById('modal-category').classList.remove('hidden');
        initializePagination('modal-category', 'li', '.pagination-info', '.pagination-controls', 6);
    });

    document.querySelector('[data-tab="user-roles-tab"]').addEventListener('click', function() {
        initializePagination('user-roles-tab', '.user-role-item', '.pagination-info-roles', '.pagination-controls-roles', 2);
    });
});