// Mobile menu management
document.addEventListener('DOMContentLoaded', function () {
    const mobileMenuButton = document.getElementById('mobile-menu-button');
    const mobileMenu = document.getElementById('mobile-menu');

    mobileMenuButton.addEventListener('click', function () {
        mobileMenu.classList.toggle('hidden');
    });

    // Close the menu when you click elsewhere
    document.addEventListener('click', function (event) {
        if (!mobileMenuButton.contains(event.target) &&
            !mobileMenu.contains(event.target) &&
            !mobileMenu.classList.contains('hidden')) {
            mobileMenu.classList.add('hidden');
        }
    });

    const textarea = document.getElementById('content-textarea');
    const preview = document.getElementById('markdown-preview');
    if (textarea && preview && window.markdown) {
        textarea.addEventListener('input', function() {
            preview.innerHTML = markdown.toHTML(textarea.value);
        });
        preview.innerHTML = markdown.toHTML(textarea.value);
    }
});

// Reusable function to initialize pagination
function initializePagination(containerId, itemSelector, paginationInfo, paginationControls, itemsPerPage = 4) {
    const container = document.getElementById(containerId);
    if (!container) return;

    const items = container.querySelectorAll(itemSelector);
    if (items.length === 0) {
        document.querySelector(paginationInfo).textContent = 'No items to display';
        return;
    }

    const totalItems = items.length;
    const totalPages = Math.ceil(totalItems / itemsPerPage);
    let currentPage = 1;

    // Function for displaying elements on a specific page
    function showPage(page) {
        // Hide all elements
        items.forEach(item => {
            item.classList.add('hidden');
        });

        // Calculate start and end indices for the current page
        const startIndex = (page - 1) * itemsPerPage;
        const endIndex = Math.min(startIndex + itemsPerPage, totalItems);

        // Show items on current page
        for (let i = startIndex; i < endIndex; i++) {
            items[i].classList.remove('hidden');
        }

        // Update page indicator
        const infoElement = container.querySelector(paginationInfo);
        if (infoElement) {
            infoElement.textContent = `Page ${page}/${totalPages}`;
        }

        // Update pagination buttons
        const buttons = container.querySelectorAll(`${paginationControls} button[data-page]`);
        buttons.forEach(button => {
            if (parseInt(button.dataset.page) === page) {
                button.classList.remove('bg-gray-200', 'hover:bg-gray-300', 'text-gray-800');
                button.classList.add('bg-blue-600', 'text-white');
            } else {
                button.classList.remove('bg-blue-600', 'text-white');
                button.classList.add('bg-gray-200', 'hover:bg-gray-300', 'text-gray-800');
            }
        });

        // Update previous/next button status
        const prevButton = container.querySelector(`${paginationControls} button[data-action="prev"]`);
        const nextButton = container.querySelector(`${paginationControls} button[data-action="next"]`);

        if (prevButton) {
            prevButton.disabled = page === 1;
            if (page === 1) {
                prevButton.classList.add('opacity-50', 'cursor-not-allowed');
            } else {
                prevButton.classList.remove('opacity-50', 'cursor-not-allowed');
            }
        }

        if (nextButton) {
            nextButton.disabled = page === totalPages;
            if (page === totalPages) {
                nextButton.classList.add('opacity-50', 'cursor-not-allowed');
            } else {
                nextButton.classList.remove('opacity-50', 'cursor-not-allowed');
            }
        }

        currentPage = page;
    }

    // Create pagination controls
    function createPaginationControls() {
        const paginationDiv = container.querySelector(paginationControls);
        if (!paginationDiv) return;

        // Empty existing controls
        paginationDiv.innerHTML = '';

        // Previous button
        const prevButton = document.createElement('button');
        prevButton.textContent = '« Précédent';
        prevButton.dataset.action = 'prev';
        prevButton.className = 'py-2 px-4 rounded bg-gray-200 hover:bg-gray-300 text-gray-800';
        prevButton.addEventListener('click', () => {
            if (currentPage > 1) {
                showPage(currentPage - 1);
            }
        });
        paginationDiv.appendChild(prevButton);

        // Page buttons
        for (let i = 1; i <= totalPages; i++) {
            const pageButton = document.createElement('button');
            pageButton.textContent = i;
            pageButton.dataset.page = i;
            pageButton.className = 'py-2 px-4 rounded bg-gray-200 hover:bg-gray-300 text-gray-800';
            pageButton.addEventListener('click', () => showPage(i));
            paginationDiv.appendChild(pageButton);
        }

        // Next button
        const nextButton = document.createElement('button');
        nextButton.textContent = 'Suivant »';
        nextButton.dataset.action = 'next';
        nextButton.className = 'py-2 px-4 rounded bg-gray-200 hover:bg-gray-300 text-gray-800';
        nextButton.addEventListener('click', () => {
            if (currentPage < totalPages) {
                showPage(currentPage + 1);
            }
        });
        paginationDiv.appendChild(nextButton);
    }

    createPaginationControls();
    showPage(1);
}