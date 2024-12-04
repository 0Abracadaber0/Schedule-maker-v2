var modal = document.getElementById("editModal");
var span = document.getElementsByClassName("close")[0];
var saveBtn = document.getElementById("saveBtn");
var subjectInput = document.getElementById("subjectInput");
var deleteBtn = document.getElementById("deleteBtn");

var currentButton;  // Переменная для хранения текущей кнопки

// Функция для открытия модального окна
function openModal(button) {
    currentButton = button;  // Получаем кнопку, в которой был клик
    var subjectName = currentButton.textContent;  // Получаем название предмета из кнопки
    subjectInput.value = subjectName;  // Вставляем название предмета в поле ввода
    modal.style.display = "flex";  // Показываем модальное окно
}

// Обработчик клика по кнопке "Редактировать"
function rebindEditButtons() {
    var editBtns = document.querySelectorAll('.editBtn');
    editBtns.forEach(function(btn) {
        btn.addEventListener('click', function() {
            openModal(this);  // Открываем модальное окно для этой кнопки
        });
    });
}

// Изначальная регистрация обработчиков
document.addEventListener('DOMContentLoaded', function() {
    rebindEditButtons();  // Регистрируем обработчики событий для кнопок редактирования
});

// Обработчик закрытия модального окна
span.onclick = function() {
    modal.style.display = "none";  // Закрываем модальное окно
    currentButton = null;  // Сбрасываем текущую кнопку
};

// Сохранение изменений
saveBtn.addEventListener('click', function() {
    var newSubjectName = subjectInput.value.trim();  // Получаем новое название предмета
    if (newSubjectName) {
        // Обновляем текст только в кнопке
        currentButton.textContent = newSubjectName;
    }
    modal.style.display = "none";  // Закрываем модальное окно
    currentButton = null;  // Сбрасываем текущую кнопку

    // Перерегистрируем обработчики событий для кнопок редактирования после сохранения
    rebindEditButtons();
});

// Удаление строки
deleteBtn.addEventListener('click', function() {
    if (currentButton) {
        var row = currentButton.closest('tr');  // Находим строку, содержащую кнопку
        row.remove();  // Удаляем строку из таблицы
    }
    modal.style.display = "none";  // Закрываем модальное окно
    currentButton = null;  // Сбрасываем текущую кнопку

    // Перерегистрируем обработчики событий для кнопок редактирования после удаления
    rebindEditButtons();
});

// Закрытие модального окна, если кликнули вне его
window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";  // Закрываем модальное окно
        currentButton = null;  // Сбрасываем текущую кнопку
    }
};
