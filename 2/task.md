# Задача о посадке пассажиров

## Описание задачи

Вы работаете на авиакомпанию и вам необходимо решить задачу о посадке пассажиров в самолет. Самолет имеет один ряд с шестью местами: A, B, C, D, E и F. Пассажиры поступают в очереди и должны занять свои места в самолете.

При посадке пассажиры занимают свои места по следующим правилам:
- Пассажиры могут занимать места только в порядке A, B, C, D, E, F.
- Если пассажир не может занять свое место, потому что оно уже занято, он блокирует проход и ждет, пока его место освободится.
- Когда место пассажира освобождается, следующий пассажир в очереди может занять это место.

Вам нужно написать программу, которая принимает на вход количество пассажиров и последовательность их появления в очереди, а затем определяет, сколько времени (единиц времени) потребуется, чтобы все пассажиры заняли свои места в самолете.

## Формат ввода

Первая строка содержит целое число n (1 ≤ n ≤ 100) - количество пассажиров.

Следующие n строк содержат информацию о пассажирах, каждая строка имеет следующий формат:

`t seat`

где t - время прибытия пассажира в очереди (0 ≤ t ≤ 1000), seat - место, которое пассажир пытается занять (A, B, C, D, E или F).

## Формат вывода

Одно целое число, равное количеству единиц времени, необходимых для посадки всех пассажиров.

## Примеры

### Ввод
```
6
0 1A
0 1B
0 1C
0 1D
0 1E
0 1F
```

### Вывод
```
26
```

### Ввод
```
6
0 1A
0 2B
0 3C
0 4D
0 5E
0 6F
```

### Вывод
```
11
```

### Ввод
```
6
10 2C
20 2B
30 2A
10 1D
10 1E
10 1F
```

### Вывод
```
91
```

## Пояснения

Рассмотрим третий пример.

Момент времени 1. Проход первого ряда заблокирован первым пассажиром (2С).

Момент времени 2. Проход первого

ряда заблокирован вторым пассажиром (2B), второго – первым (2C).

Момент времени 3-12. Проход первого ряда заблокирован вторым пассажиром (2B), второго – первым (2C), который размещает вещи на полке.

Момент времени 13. Проход первого ряда заблокирован третьим пассажиром (2A), второго – вторым (2B).

Момент времени 14-33. Проход первого ряда заблокирован третьим пассажиром (2A), второго – вторым (2B), который размещает вещи на полке.

Момент времени 34-38. Проход первого ряда заблокирован третьим пассажиром (2A), второго – вторым (2B), которого пропускает на место первый пассажир (1С).

Момент времени 39. Проход первого ряда заблокирован четвертым пассажиром (1D), второго – третьим (2А).

Момент времени 40-49. Проход первого ряда заблокирован четвертым пассажиром (1D), который размещает вещи на полке, второго – третьим (2А), который размещает вещи на полке.

Момент времени 50. Проход первого ряда заблокирован пятым пассажиром (1E), второго – третьим (2А), который размещает вещи на полке.

Момент времени 51-60. Проход первого ряда заблокирован пятым пассажиром (1Е), который размещает вещи на полке, второго – третьим (2А), который размещает вещи на полке.

Момент времени 61-65. Проход первого ряда заблокирован пятым пассажиром (1Е), которого пропускает на место четвертый пассажир (1D), второго – третьим (2А), который размещает вещи на полке.

Момент времени 66. Проход первого ряда заблокирован шестым пассажиром (1F), второго – третьим (2А), который размещает вещи на полке.

Момент времени 67-69. Проход первого ряда заблокирован шестым пассажиром (1F), который размещает вещи на полке, второго – третьим (2А), который размещает вещи на полке.

Момент времени 70-76. Проход первого ряда заблокирован шест

ым пассажиром (1F), который размещает вещи на полке, второго – третьим (2А), которого пропускают на место первый (1С) и второй (1B) пассажиры.

Момент времени 77-84. Проход первого ряда заблокирован шестым пассажиром (1F), которого пропускают на место четвертый (1D) и пятый (1E) пассажиры, второго – третьим (2А), которого пропускают на место первый (2C) и второй (2B) пассажиры.

Момент времени 85-91. Проход первого ряда заблокирован шестым пассажиром (1F), которого пропускают на место четвертый (1D) и пятый (1E) пассажиры.

Таким образом, всего потребуется 91 единиц времени для посадки всех пассажиров.