function showEditForm(commentId) {
    document.getElementById('comment-content-' + commentId).classList.add('hidden');
    document.getElementById('edit-comment-btn-' + commentId).classList.add('hidden');
    document.getElementById('edit-comment-form-' + commentId).classList.remove('hidden');
}

function hideEditForm(commentId) {
    document.getElementById('comment-content-' + commentId).classList.remove('hidden');
    document.getElementById('edit-comment-btn-' + commentId).classList.remove('hidden');
    document.getElementById('edit-comment-form-' + commentId).classList.add('hidden');
}