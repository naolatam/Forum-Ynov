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

document.getElementById('open-delete-modal').onclick = function () {
    document.getElementById('delete-modal').classList.remove('hidden');
};
document.getElementById('cancel-delete').onclick = function () {
    document.getElementById('delete-modal').classList.add('hidden');
};