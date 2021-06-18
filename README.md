# trueconf-test
## Небольшое пояснение к решению
1. В задаче не указано что идентификатор пользователя должен быть уникальным, поэтому уникальность не реализована и пользователь сам вводит идентификатор при добавление. Уникальность можно было бы добавить с помощью счётчика и инкрементировать счётчик после добавления нового объекта, а ввод идентификатора пользователя игнорировать при добавление.
2. Для каждого объекта можно было бы создавать новый json-файл, но я решил что это будет сложнее реализовывать, поэтому все данные пишутся в один файл. Файл полностью перезаписывается после каждой изменяющей базу операции.
## Условия задачи
Написать маленькое приложение на ECHO framework предоставляющие REST API по работе с сущностью User.

REST API должно удовлетворять следующие возможности:
* Добавление User
* Получение списка User
* Получение User по Id
* Редактирование User по Id
* Удаление User по Id

REST API должно работать с форматом данных JSON.

Сущность User должно состоять минимум из следующих полей:
* Идентификатор пользователя
* Отображаемое имя

Можно использовать дополнительные поля, если считаете нужным.

В качестве хранилища данных необходимо использовать файл в формате JSON.

