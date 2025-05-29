document.addEventListener('DOMContentLoaded', function() {
    const openModalButton = document.querySelector('#edit-profile');
    const modal = document.getElementById('profile-modal');
    const closeModalButton = document.getElementById('close-modal');
    const cancelModalButton = document.getElementById('cancel-modal');
    const modalBackdrop = document.getElementById('modal-backdrop');
    const deleteButton = document.getElementById('delete-account');
    const confirmDeleteModal = document.getElementById('confirm-delete-modal');
    const cancelDeleteButton = document.getElementById('cancel-delete');
    const confirmDeleteButton = document.getElementById('confirm-delete');

    openModalButton.addEventListener('click', function(e) {
        e.preventDefault();
        modal.classList.remove('hidden');
        document.body.classList.add('overflow-hidden');
    });

    const closeModal = function() {
        modal.classList.add('hidden');
        document.body.classList.remove('overflow-hidden');
    };

    deleteButton.addEventListener('click', function() {
        confirmDeleteModal.classList.remove('hidden');
    });

    cancelDeleteButton.addEventListener('click', function() {
        confirmDeleteModal.classList.add('hidden');
    });

    confirmDeleteButton.addEventListener('click', function() {
        // logique pour supprimer le compte
        confirmDeleteModal.classList.add('hidden');
        modal.classList.add('hidden');
    });

    closeModalButton.addEventListener('click', closeModal);
    cancelModalButton.addEventListener('click', closeModal);
    modalBackdrop.addEventListener('click', closeModal);
});