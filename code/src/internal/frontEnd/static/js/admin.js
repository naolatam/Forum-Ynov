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

// Gestion des tabs
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