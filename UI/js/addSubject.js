// Получаем элементы
var addSubjectBtn = document.getElementById("addSubjectBtn");  // Кнопка "Добавить предмет"
var addSubjectModal = document.getElementById("addSubjectModal");  // Модальное окно для добавления
var newSubjectInput = document.getElementById("newSubjectInput");  // Поле ввода
var addSubjectBtnModal = document.getElementById("addSubjectBtnModal");  // Кнопка "Добавить" в модальном окне
var closeModal = document.querySelector(".close");  // Кнопка закрытия модального окна

// Модальное окно для редактирования
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

// Открытие модального окна для добавления нового предмета
addSubjectBtn.addEventListener('click', function() {
    addSubjectModal.style.display = "flex";  // Показываем модальное окно
    newSubjectInput.value = '';  // Очищаем поле ввода
});

// Закрытие модального окна при нажатии на кнопку "закрыть"
closeModal.addEventListener('click', function() {
    addSubjectModal.style.display = "none";  // Закрываем модальное окно
});

// Обработчик для кнопки "Добавить" в модальном окне
addSubjectBtnModal.addEventListener('click', function() {
    var newSubject = newSubjectInput.value.trim();  // Получаем введенное название

    if (newSubject) {
        // Добавляем новую строку в таблицу
        var tableBody = document.querySelector("#subjectsTable tbody");
        var newRow = document.createElement("tr");
        var newCell = document.createElement("td");
        var newButton = document.createElement("button");
        
        newButton.textContent = newSubject;
        newButton.classList.add("editBtn");
        
        newCell.appendChild(newButton);
        newRow.appendChild(newCell);
        tableBody.appendChild(newRow);  // Добавляем новую строку в таблицу

        // Перерегистрируем обработчики для новой кнопки
        rebindEditButtons();
    }

    addSubjectModal.style.display = "none";  // Закрываем модальное окно после добавления
});

// Обработчик закрытия модального окна
var spanAdd = document.getElementsByClassName("close")[1]; // Кнопка закрытия для добавления
spanAdd.onclick = function() {
    addSubjectModal.style.display = "none";
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
