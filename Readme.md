#Напоминалки на устные вопросы
##1. Самый эффективный способ слияния строк является copy()
##2. Это методы, определяющие поведение разных обьектов
##3. RWMutex дает расширенный функционал, позволяющий параллельное чтение данных. Но обычный mutex быстрее работает
##4. Буферизированные каналы не блокирует горутину, пока не будет заполнен. Небуф. канал всегда блокирует пока не прочитано или не записано.
##5. 0
##6. Перегрузки
##7. По порядку
##8. new() выделяет память и возвращает указатель на область памяти, make() выделяет место в памяти только для слайсов, мап и каналов, создает обьект
##9. var, make - для map, var, make, := - для слайс
##10 Мы переопределяем в функции и они изменятся только там
##11 В горутину добавляют копию wg
##12 В условии переопределяется n и все что с n происходит внутри if
##13 В someAction когда изменяется на 100 и оно работает из-за того,что это тот же слайс, а в append уже v это ссылка на другой массив, который мы не возвращаем из функции
##14 Такая же причина, что и в 13, работаем с разными массивами