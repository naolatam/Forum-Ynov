document.addEventListener('DOMContentLoaded', function () {
    const select = document.getElementById('tag-select');
    const tagsContainer = document.getElementById('selected-tags');

    function updateTags() {
        tagsContainer.innerHTML = '';
        Array.from(select.selectedOptions).forEach(option => {
            const span = document.createElement('span');
            span.className = 'bg-blue-600 text-white px-3 py-1 rounded-full text-sm flex items-center';
            span.innerHTML = `
                ${option.text}
                <button class="ml-2 focus:outline-none" data-value="${option.value}">
                    <i class="fa-solid fa-xmark"></i>
                </button>
            `;
            tagsContainer.appendChild(span);
        });
    }

    select.addEventListener('change', updateTags);

    tagsContainer.addEventListener('click', function (e) {
        if (e.target.closest('button')) {
            const value = e.target.closest('button').getAttribute('data-value');
            Array.from(select.options).forEach(opt => {
                if (opt.value === value) opt.selected = false;
            });
            updateTags();
        }
    });

    updateTags();
});

document.addEventListener('DOMContentLoaded', function () {
    const openDeleteModalBtn = document.getElementById('open-delete-modal');
    const deleteModal = document.getElementById('delete-modal');
    const cancelDeleteBtn = document.getElementById('cancel-delete');

    if (openDeleteModalBtn && deleteModal) {
        openDeleteModalBtn.onclick = function () {
            deleteModal.classList.remove('hidden');
        };
    }
    if (cancelDeleteBtn && deleteModal) {
        cancelDeleteBtn.onclick = function () {
            deleteModal.classList.add('hidden');
        };
    }
});

document.addEventListener('DOMContentLoaded', function () {
    const imageInput = document.getElementById('image');
    const imagePreview = document.getElementById('post-image-preview');
    const errorModal = document.getElementById('post-image-error-modal');
    const errorCloseBtn = document.getElementById('post-image-error-close');

    if (imageInput && imagePreview && errorModal && errorCloseBtn) {
        imageInput.addEventListener('change', (e) => {
            const file = e.target.files[0];
            if (file) {
                if (file.size > 20 * 1024 * 1024) {
                    errorModal.classList.remove('hidden');
                    imageInput.value = '';
                } else {
                    const reader = new FileReader();
                    reader.onload = function (e) {
                        imagePreview.src = e.target.result;
                    }
                    reader.readAsDataURL(file);
                }
            }
        });

        errorCloseBtn.addEventListener('click', () => {
            errorModal.classList.add('hidden');
            imageInput.value = '';
        });
    }
});