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

document.addEventListener('DOMContentLoaded', function() {
    const editAvatarBtn = document.querySelector('.relative.mb-3 button');
    const avatarUploadModal = document.getElementById('avatar-upload-modal');
    const closeAvatarModalBtn = document.getElementById('close-avatar-modal');
    const cancelAvatarUploadBtn = document.getElementById('cancel-avatar-upload');
    const avatarInput = document.getElementById('avatar-upload');
    const avatarPreview = document.getElementById('avatar-preview');

    // Open Avatar Upload Modal
    if (editAvatarBtn) {
        editAvatarBtn.addEventListener('click', (e) => {
            e.preventDefault();
            avatarUploadModal.classList.remove('hidden');
        });
    }

    function closeAvatarModal() {
        avatarUploadModal.classList.add('hidden');
    }

    if (closeAvatarModalBtn) closeAvatarModalBtn.addEventListener('click', closeAvatarModal);
    if (cancelAvatarUploadBtn) cancelAvatarUploadBtn.addEventListener('click', closeAvatarModal);

    // Image Preview
    if (avatarInput) {
        avatarInput.addEventListener('change', (e) => {
            const file = e.target.files[0];
            if (file) {
                const reader = new FileReader();
                reader.onload = function(e) {
                    avatarPreview.src = e.target.result;
                }
                reader.readAsDataURL(file);
            }
        });
    }

    const confirmUploadBtn = document.getElementById('confirm-avatar-upload');
    if (confirmUploadBtn) {
        confirmUploadBtn.addEventListener('click', () => {
            const newAvatarSrc = avatarPreview.src;
            document.querySelectorAll('.profile-avatar').forEach(img => {
                img.src = newAvatarSrc;
            });

            closeAvatarModal();
        });
    }
});

document.addEventListener('DOMContentLoaded', function() {
    initializePagination('content-tab', '.content-item', '.pagination-info-content', '.pagination-controls-content', 4);
});