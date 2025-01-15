1) 1. Two Sum
	Линк: https://leetcode.com/problems/two-sum/
	Сколько раз решал: 2
	Время: 08:04
	Комментарий: Внимательней с использованием таблицы. Была синтаксическая ошибка в получении значения из мапы.
	Tag: HashMap
2) 205. Isomorphic Strings
	Линк: https://leetcode.com/problems/isomorphic-strings/
	Сколько раз решал: 1
	Время: 22:14
	Комментарий: Внимательно с возвращаемым значением. Не забывай, что нужно две мапы, а не одна.
				 Так как есть такой кейс: s="badc" t="baba"
	Tag: HashMap
3) 36. Valid Sudoku
	Линк: https://leetcode.com/problems/valid-sudoku/
	Сколько раз решал: 1
	Время: 24:43
	Комментарий: Будь внимателен! Не забывай добавлять элементы в мапу после проверки!!!
	Tag: HashMap
4) 347. Top K Frequent Elements
	Линк: https://leetcode.com/problems/top-k-frequent-elements/
	Сколько раз решал: 1
	Время: 19:55
	Комментарий: Внимательно со знаками сравнения, когда работаешь с обратным счетчиком.
	Tag: HashMap
5) 242. Valid Anagram
	Линк: https://leetcode.com/problems/valid-anagram/
	Сколько раз решал: 1
	Время: 06:41
	Комментарий: Решилось легко.
	Tag: HashMap
6) 49. Group Anagrams
	Линк: https://leetcode.com/problems/group-anagrams/
	Сколько раз решал: 1
	Время: 10:50
	Комментарий: Используй код строки(количество тех или иных букв) как ключ для мапы, а значение массив сткрок,
				 подходящих под ключ.
	Tag: HashMap
7) 2657. Find the Prefix Common Array of Two Arrays
	Линк: https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
	Сколько раз решал: 1
	Время: 20:20
	Комментарий: Стоит запомнить, что два вхождения означают, что число уже было. Это сильно поможет в решении.
	Tag: HashMap
8) 274. H-Index
	Линк: https://leetcode.com/problems/h-index/
	Сколько раз решал: 1
	Время: 15:01
	Комментарий: Тяжело идет. Решать не через сортировку!
	TODO: Прорешать несколько раз.
	Tag: HashMap
9) 266. Palindrome Permutation
	Линк: https://www.lintcode.com/problem/916/
	Сколько раз решал: 1
	Время: 08:14
	Комментарий: Внимательно с последним условием!!! Значение должно быть меньше или равно 1.
	Tag: HashMap
10) 356. Line Reflection
	Линк: https://www.lintcode.com/problem/908/
	Сколько раз решал: 1
	Время: 28:28
	Комментарий: Нужно решить самому без просмотра видео. Помнить о формуле minX + maxX - x.
				 Объяснить появление формулы. На lintcode может приходить пустой массив.
	TODO: Прорешать несколько раз.
	Tag: HashMap
11) 1464. Maximum Product of Two Elements in an Array
	Линк: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
	Сколько раз решал: 1
	Время: 07:45
	Комментарий: Внимательно с возвращаемым значением, вернуть НЕ индексы, а результат.
	Tag: IterateArray
	Interview: 1
12) 896. Monotonic Array
	Линк: https://leetcode.com/problems/monotonic-array/
	Сколько раз решал: 1
	Время: 07:00
	Комментарий: Все ок.
	Tag: IterateArray
	Interview: 1
13) 303. Range Sum Query - Immutable
	Линк: https://leetcode.com/problems/range-sum-query-immutable/
	Сколько раз решал: 2
	Время: 10:26
	Комментарий: Внимательно с методом, передается ссылка на объект, в котором префиксный массив, а не сам префиксный
				 массив.
	Tag: PrefixArray
14) 724. Find Pivot Index
	Линк: https://leetcode.com/problems/find-pivot-index/
	Сколько раз решал: 1
	Время: 28:45
	Комментарий: Решить самому без просмотра видео.
	Подсказка: нужна сумма всех элементов, левая и правая сумма.
	Tag: AllElementsSum
15) 560. Subarray Sum Equals K
	Линк: https://leetcode.com/problems/subarray-sum-equals-k/description/
	Сколько раз решал: 1
	Время: очень долго
	Комментарий: Решить самому без просмотра видео.
	Подсказка: Мапа хранит префиксные суммы. Текущая сумма элементов минус k - это количство поддмассивов с суммой k,
			   которые можно составить.
	Tag: PrefixArray
16) 268. Missing Number
	Линк: https://leetcode.com/problems/missing-number/
	Сколько раз решал: 2
	Время: 03:30
	Комментарий: Вниметельно со значениями инициализации, так как ожидаемое значение должно быть проинициализировано
				 длинной массива.
	Tag: AllElementsSum
17) 189. Rotate Array
	Линк: https://leetcode.com/problems/rotate-array/
	Сколько раз решал: 1
	Время: 10:46
	Комментарий: Необходимо самому объяснить почему работает и используется k = k % len(nums).
	Tag: SplitArray
18) 674. Longest Continuous Increasing Subsequence
	Линк: https://leetcode.com/problems/longest-continuous-increasing-subsequence/
	Сколько раз решал: 1
	Время: 06:12
	Комментарий: Нужно вернуть максимальную длину, а не последнюю. Не забыть проверку.
	Tag: FindMax
19) 1984. Minimum Difference Between Highest and Lowest of K Scores
	Линк: https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
	Сколько раз решал: 1
	Время: 08:28
	Комментарий: Решить без видео.
	TODO: Прорешать несколько раз.
	Tag: SortArray
20) 1310. XOR Queries of a Subarray
	Линк: https://leetcode.com/problems/xor-queries-of-a-subarray/
	Сколько раз решал: 1
	Время: 08:28
	Комментарий: XOR (go ^), обратная операция XOR - это XOR. Создаем префиксный массив, где первый элемент оставляем 0,
				 а остальные как обычно, результат следущего равен текущий из префикса XOR текущий из входного. Дальше
				 находим "разницу" (в этом случае XOR) для каждого запроса, но помним, что правый всегда берем следующий
				 за указанным в запросе. То есть запрос [0, 1] значит мы в префиксе берем [0, 2].
	Tag: PrefixArray
21) 304. Range Sum Query 2D - Immutable
	Линк: https://leetcode.com/problems/range-sum-query-2d-immutable/
	Сколько раз решал: 1
	Время: 98:28
	Комментарий: Составляем префиксную матрицу размером на 1 больше, чтобы первые строка и колонка были равны 0.
				 Дальше нужно рисовать. Формула по созданию префиксной матрицы:
				 prefix[i][j] = matrix[i - 1][j - 1] + prefix[i - 1][j] + prefix[i][j - 1] - prefix[i - 1][j - 1]
				 ФОрмула для получения значения:
				 fullSum := this.prefixMatrix[row2 + 1][col2 + 1]
				 above := this.prefixMatrix[row1][col2 + 1]
				 left := this.prefixMatrix[row2 + 1][col1]
				 result = fullSum - above - left + this.prefixMatrix[row1][col1]
	Tag: PrefixArray
22) 17. Letter Combinations of a Phone Number
	Линк: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
	Сколько раз решал: 1
	Время: 29:37
	Комментарий: Полный обход в рекурсии.
	Tag: Backtracking
23) 22. Generate Parentheses
	Линк: https://leetcode.com/problems/generate-parentheses/
	Сколько раз решал: 1
	Время: 89:35
	Комментарий: Полный обход в рекурсии, как в предыдущем случае. Используем balance для определения, что сторока
				 подходит. balance < 0 - число закрывающих больше чем открывающих - последовательность не корректна.
				 balance > n - len(currParentheses) - не хватит закрывающих скобок, чтобы сделать последовательность
				 корректной. При этих условиях выходим из функции. Если balance == 0 && len(currParentheses) == n,
				 тогда добавляем строку в ответ и выходим из функции.
	Tag: Backtracking
24) 977. Squares of a Sorted Array
	Линк: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
	Сколько раз решал: 1
	Время: 09:18
	Комментарий: Много мелких ошибок. Лучше проверять перед запуском. Ответ формируем с конца.
				 Не забываем менять указатели.
	Tag: TwoPointers
25) 125. Valid Palindrome
	Линк: https://leetcode.com/problems/valid-palindrome/
	Сколько раз решал: 1
	Время: 10:35
	Комментарий: Много мелких ошибок. Лучше проверять перед запуском. Не забываем про приведение к Lower case.
				 Не забываем менять указатели.
	Tag: TwoPointers
26) 167. Two Sum II - Input Array Is Sorted
	Линк: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
	Сколько раз решал: 1
	Время: 04:24
	Комментарий: Ok.
	Tag: TwoPointers
27) 15. 3Sum
	Линк: https://leetcode.com/problems/3sum/
	Сколько раз решал: 1
	Время: 22:17
	Комментарий: Сортировка и фиксирование одного из трех элементов. t = m = O(n^2)
	Tag: TwoPointers
28) 16. 3Sum Closest
	Линк: https://leetcode.com/problems/3sum-closest/
	Сколько раз решал: 1
	Время: 30:42
	Комментарий: Сортировка и фиксирование одного из трех элементов. Не забыть сравнить target - cur.
	Tag: TwoPointers
29) 11. Container With Most Water
	Линк: https://leetcode.com/problems/container-with-most-water/
	Сколько раз решал: 1
	Время: 05:53
	Комментарий: Не торопиться.
	Tag: TwoPointers
30) Intersection Of Sorted Arrays
	Линк: https://www.interviewbit.com/problems/intersection-of-sorted-arrays/
	Сколько раз решал: 1
	Время: 06:11
	Комментарий: Не торопиться.
	Tag: TwoPointersOnTwoArrays
31) 844. Backspace String Compare
	Линк: https://leetcode.com/problems/backspace-string-compare/
	Сколько раз решал: 1
	Время: 44:39
	Комментарий: Идея понятно, но идет очень тяжело. Не забыть сравнить в конце.
	Tag: TwoPointersOnTwoArrays
32) 283. Move Zeroes
	Линк: https://leetcode.com/problems/move-zeroes/
	Сколько раз решал: 1
	Время: 10:15
	Комментарий: Не торопиться.
	Tag: FastSlowPointer
33) 14. Longest Common Prefix
	Линк: https://leetcode.com/problems/longest-common-prefix/
	Сколько раз решал: 1
	Время: 08:33
	Комментарий: Не забыть проверить во вложенном цикле, что индекс может быть за границей строки.
	Tag: TwoPointersOnTwoArrays
34) One Edit Distance
	Линк: https://www.lintcode.com/problem/640/
	Сколько раз решал: 1
	Время: 38:33
	Комментарий: Как только встречаем отличающийся символ, то проверяем все оставшиеся варианты:
				 return s[sIdx+1:] == t[tIdx:] || s[sIdx:] == t[tIdx+1:] || s[sIdx+1:] == t[tIdx+1:]. Не забыть
				 проверить после цикла, что len(s)-len(t) != 0, так как если закончили цикл и строки равны, то это
				 значит false, так как они идентичны.
	Tag: TwoPointersOnTwoArrays
35) 228. Summary Ranges
	Линк: https://leetcode.com/problems/summary-ranges/
	Сколько раз решал: 1
	Время: 10:01
	Комментарий: быть внимательней.
	Tag:  Non-overlapping SlidingWindows
36) 485. Max Consecutive Ones
	Линк: https://leetcode.com/problems/max-consecutive-ones/
	Сколько раз решал: 1
	Время: 11:25
	Комментарий: сначала ищем все одинаковые послеовательности, а потом разбираемся 1 это или 0.
	Tag:  Non-overlapping SlidingWindows
37) 849. Maximize Distance to Closest Person
	Линк: https://leetcode.com/problems/maximize-distance-to-closest-person/
	Сколько раз решал: 1
	Время: 32:14
	Комментарий: Отвратительная задача. Пустые места считаются как промежутки между цифрами, а не сами места. Стоит
				 нарисовать крайние случае, когда нет никого слева или справа. Левый указатель выражем через правый.
				 Правый двигаем в отдельном вложенном цикле пока не отличабтся значения.
	Tag:  Non-overlapping SlidingWindows
38) 443. String Compression
	Линк: https://leetcode.com/problems/maximize-distance-to-closest-person/
	Сколько раз решал: 1
	Время: 31:00
	Комментарий: Можно обойтись одним условием num > 1, а присвоение буквы вынести в начало.
	Tag:  Non-overlapping SlidingWindows
39) 1004. Max Consecutive Ones III
	Линк: https://leetcode.com/problems/max-consecutive-ones-iii/
	Сколько раз решал: 1
	Время: 30:42
	Комментарий: Если начинаешь с нуля оба указателя, то растояние будте right - left без добавления 1,
				 так как в отличие от пересекающихся окон, правый будет уходить правее нужного окна.
	Tag: Overlapping SlidingWindows
40) 3. Longest Substring Without Repeating Characters
	Линк: https://leetcode.com/problems/longest-substring-without-repeating-characters/
	Сколько раз решал: 1
	Время: 17:45
	Комментарий: Правый начинает с -1, так как изначально мапа с символами пуста, если 0, то нужно сразу класть первый
				 символ.
	Tag: Overlapping SlidingWindows
41) 567. Permutation in String
	Линк: https://leetcode.com/problems/permutation-in-string/
	Сколько раз решал: 1
	Время: 73:55
	Комментарий: Очень тяжело. Правый двигаем только внутри вложенного цикла. В мапу второй строки добавляем все символы
				 посредством уменьшения количества вхождений, когда двигаем левую границу. Таким образом у нас
				 получается, что 0 - 1 = -1 и в условии вложеннго цикла это позволит подвинуть правую границу.
	Tag: Overlapping SlidingWindows
42) 704. Binary Search
	Линк: https://leetcode.com/problems/binary-search/
	Сколько раз решал: 1
	Время: 09:22
	Комментарий: Меняется только функция определения подходит элемент или нет.
	Tag: BinarySearch
43) 34. Find First and Last Position of Element in Sorted Array
	Линк: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
	Сколько раз решал: 1
	Время: 20:44
	Комментарий: 2 байнари серча, сначала ищем конец, потом начало.
	Tag: BinarySearch
44) 74. Search a 2D Matrix
	Линк: https://leetcode.com/problems/search-a-2d-matrix/
	Сколько раз решал: 2
	Время: 05:49
	Комментарий: 2 байнари серча, сначала ищем строку, потом таргет. Или превращаем все в один массив по средством
				 вычисления указателей -> i = cur/n, j = cur%n, где cur - это индекс псевдо одномерного массива, а
				 n длина строки len(matrix[0]).
	Tag: BinarySearch
45) 33. Search in Rotated Sorted Array
	Линк: https://leetcode.com/problems/search-in-rotated-sorted-array/
	Сколько раз решал: 1
	Время: 17:13
	Комментарий: Важно правильно выбрать good. Здесь лучше nums[mid] > nums[len(nums) - 1].
	Tag: BinarySearch
46) 69. Sqrt(x)
	Линк: https://leetcode.com/problems/sqrtx/
	Сколько раз решал: 1
	Время: 05:02
	Комментарий: Ищем среди квадратов, то есть i*i.
	Tag: BinarySearch
47) 153. Find Minimum in Rotated Sorted Array
	Линк: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
	Сколько раз решал: 1
	Время: 05:25
	Комментарий: Ответ в правом итераторе.
	Tag: BinarySearch
48) 367. Valid Perfect Square
	Линк: https://leetcode.com/problems/valid-perfect-square/
	Сколько раз решал: 1
	Время: 06:08
	Комментарий: Ищем среди квадратов, то есть mid*mid.
	Tag: BinarySearch
49) 852. Peak Index in a Mountain Array
	Линк: https://leetcode.com/problems/peak-index-in-a-mountain-array/
	Сколько раз решал: 1
	Время: 06:08
	Комментарий: Края не могу быть ответом, функция текущий меньше предыдущего.
	Tag: BinarySearch
50) 81. Search in Rotated Sorted Array II
	Линк: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
	Сколько раз решал: 1
	Время: 67:13
	Комментарий: Совсем не ложиться решение. Проверяем, что серединный не таргет. Избавляемся от одинаковых элементов,
				 то есть идем слева пока соседние элементы не станут отличаться. Тоже самое справа. Потом определяем
				 какая часть отсортирована. nums[mid] >= nums[left] // Левая половина отсортирована
				 else //Правая половина отсортирована. Внутри каждой ветки проверяем где таргет и сужаем диапазон.
	Tag: BinarySearch
51) 206. Reverse Linked List
	Линк: https://leetcode.com/problems/reverse-linked-list/
	Сколько раз решал: 1
	Время: 08:46
	Комментарий: Просто меняем направление указателей.
	Tag: LinkedList
52) 876. Middle of the Linked List
	Линк: https://leetcode.com/problems/middle-of-the-linked-list/
	Сколько раз решал: 1
	Время: 06:23
	Комментарий: Быстрый и медленный указатель, начинают оба с head.
	Tag: LinkedList
53) 234. Palindrome Linked List
	Линк: https://leetcode.com/problems/palindrome-linked-list/
	Сколько раз решал: 1
	Время: 09:22
	Комментарий: Решается через Middle и Reverse
	Tag: LinkedList
54) 21. Merge Two Sorted Lists
	Линк: https://leetcode.com/problems/merge-two-sorted-lists/
	Сколько раз решал: 1
	Время: 09:22
	Комментарий: Используй Dummy и Stab ноды.
	Tag: LinkedList
55) 19. Remove Nth Node From End of List
	Линк: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
	Сколько раз решал: 2
	Время: 14:42
	Комментарий: Используй Dummy ноду. Сначала идем быстрым указателем до тех N раз, важно!!! i <= N, так как Dummy ноды
				 увеличивает количество. Дальше запускаем цикл пока fast != nil и двигаем оба указателя. После завершения
				 цикла работаем только с slow, то есть slow.Next = slow.Next.Next. Мы не выйдем за пределы, так как
				 используем Dummy.
	Tag: LinkedList, TwoPointers
56) 144. Binary Tree Preorder Traversal
	Линк: https://leetcode.com/problems/binary-tree-preorder-traversal/
	Сколько раз решал: 1
	Время: 05:42
	Комментарий: Используй замыкания.
	Tag: BinaryTreeTraversal
57) 94. Binary Tree Inorder Traversal
	Линк: https://leetcode.com/problems/binary-tree-inorder-traversal/
	Сколько раз решал: 1
	Время: 02:28
	Комментарий: Используй замыкания.
	Tag: BinaryTreeTraversal
58) 145. Binary Tree Postorder Traversal
	Линк: https://leetcode.com/problems/binary-tree-postorder-traversal/
	Сколько раз решал: 1
	Время: 02:22
	Комментарий: Используй замыкания.
	Tag: BinaryTreeTraversal
59) 102. Binary Tree Level Order Traversal
	Линк: https://leetcode.com/problems/binary-tree-level-order-traversal/
	Сколько раз решал: 1
	Время: 06:14
	Комментарий: Создаем массив, в который и будем складывать значение из ноды с каждого уровня.
	Tag: BinaryTreeTraversal
60) 103. Binary Tree Zigzag Level Order Traversal
	Линк: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
	Сколько раз решал: 1
	Время: 11:14
	Комментарий: Сначала обходим и формируем массив, а потом переварачиваем нечетные слайсы.
	Tag: BinaryTreeTraversal
61) 199. Binary Tree Right Side View
	Линк: https://leetcode.com/problems/binary-tree-right-side-view/
	Сколько раз решал: 1
	Время: 05:52
	Комментарий: Записываем в массив по индексу уровня, каждый раз будет переписываться значения на более
				 правое(позднее) значение.
	Tag: BinaryTreeTraversal
62) 101. Symmetric Tree
	Линк: https://leetcode.com/problems/symmetric-tree/
	Сколько раз решал: 1
	Время: 09:18
	Комментарий: Необходимо сделать функцию, которая принимает левую ветвь и правую и сравнивает их.
	Tag: BinaryTreeTraversal
63) 100. Same Tree
	Линк: https://leetcode.com/problems/same-tree/
	Сколько раз решал: 1
	Время: 03:25
	Комментарий: Обходим деревья одновременно и сравниваем.
	Tag: BinaryTreeTraversal
64) 112. Path Sum
	Линк: https://leetcode.com/problems/path-sum/
	Сколько раз решал: 2
	Время:11:04
	Комментарий: Обходим деревья и считаем занчения. Важно доходить всегда до конца (то есть до листа), потому что
				 нода может иметь значение 0.
	Tag: BinaryTreeTraversal
65) 83. Remove Duplicates from Sorted List
	Линк: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
	Сколько раз решал: 1
	Время:06:44
	Комментарий: Обходим лист с двумя указателями(один впереди, другой сзади) и сравниваем.
	Tag: LinkedList, TwoPointers
66) 23. Merge k Sorted Lists
	Линк: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
	Сколько раз решал: 1
	Время:54:16
	Комментарий: Используем кучу. Добавляем все первые элементы листов в кучу на минимум. Далее запускаем цикл, пока
				 куча не пустая, вынимаем минимум и добавляем следующий элемент из списка после вынутого.
	Tag: Heap
67) 20. Valid Parentheses
	Линк: https://leetcode.com/problems/valid-parentheses/
	Сколько раз решал: 1
	Время: 12:05
	Комментарий: Пиши внимательно и не торопясь. Обрати внимание на последнюю проверку, в мапе с парами ищем скобку из
				 стека, а не из строки, так как из строки закрывающая.
	Tag: Stack
68) 20. Valid Parentheses
	Линк: https://leetcode.com/problems/daily-temperatures/
	Сколько раз решал: 2
	Время: 19:46
	Комментарий: Создаем стек и проходимся по массиву температур. Если стек пустой или последний элемент стека имеет
				 температуру не ниже текущей, то кладем в стек массив из двух элементов [индекс дня, температура]. Если
				 же стек не пустой и текущая температура выше последней в стеке, то тогда наступило потепление. Значит
				 начинаем во вложенном цикле вынимать элементы из стека, заполняя выходной массив разницей между текущим
				 индексом (так как это день, когда стало теплее) минус индекс из элемента стека (так как это день, для
				 которого мы хотим посчитать кол-во дней до потепления). Продолжаем пока условие "Если стек пустой или
				 последний элемент стека имеет температуру не ниже текущей" выполняется.
	Tag: Stack
69) 1249. Minimum Remove to Make Valid Parentheses
	Линк: https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
	Сколько раз решал: 1
	Время: 45:31
	Комментарий: Используй стек и пометку # для плохих скобок. Можно в конеце использовать strings.Builder.
	Tag: Stack
70) 150. Evaluate Reverse Polish Notation
	Линк: https://leetcode.com/problems/evaluate-reverse-polish-notation/
	Сколько раз решал: 1
	Время: 22:59
	Комментарий: strconv.Atoi.
	Tag: Stack
71) Balanced Parantheses!
	Линк: https://www.interviewbit.com/problems/balanced-parantheses/
	Сколько раз решал: 1
	Время: 05:32
	Комментарий: Решаем без стека. Если баланс меньше 0, то сразу выходим.
	Tag: Balance
72) 71. Simplify Path
	Линк: https://leetcode.com/problems/simplify-path/
	Сколько раз решал: 1
	Время: 25:45
	Комментарий: Используй strings.Join. strings.Split, создает пустые строки, если несколько слешей подряд.
	Tag: Stack
73) 56. Merge Intervals
	Линк: https://leetcode.com/problems/merge-intervals/
	Сколько раз решал: 1
	Время: 18:05
	Комментарий: sort.Slice(arr, func)
	Tag: Intervals
74) 986. Interval List Intersections
	Линк: https://leetcode.com/problems/interval-list-intersections/
	Сколько раз решал: 1
	Время: 22:26
	Комментарий: Всегда двигаем итератор, где меньше правая граница, даже когда интервалы не пересекаются.
	Tag: Intervals
75) 452. Minimum Number of Arrows to Burst Balloons
	Линк: https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
	Сколько раз решал: 1
	Время: 09:29:48
	Комментарий: Либо начальные точки минимальные, либо берешь первую, но тогда каунт инициализируется 1.
	Tag: Intervals
76) 253. Meeting rooms
	Линк: https://www.interviewbit.com/problems/meeting-rooms/
	Сколько раз решал: 1
	Время: 19:00
	Комментарий: Не чистые интервалы, а метод точек. Каждой точке в начале интервала даем +1, в конце -1, сортируем
				 и находим максимум. Важно, что у одинаковых точек, сначала -1, потом +1 или наоборот. То есть нужно
				 уточнить, сначала освобождается переговорка и заходят люди или наоборот.
	Tag: PointsMethod
77) 1094. Car Pooling
	Линк: https://leetcode.com/problems/car-pooling/
	Сколько раз решал: 1
	Время: 11:04
	Комментарий: Не чистые интервалы, а метод точек. Если используешь capacity, а не отдельную переменную, то сначала
				 прибавляешь (пассажиры выходят, места появляются), а потом вычитаешь.
	Tag: PointsMethod
78) 394. Decode String
	Линк: https://leetcode.com/problems/decode-string/
	Сколько раз решал: 1
	Время: 01:04:18
	Комментарий: Сложно. Проходимся по строке и складываем в стек, пока не ]. Как только встреаем, то начинаем
				 обрабатывать стек. А именно собирать подстроку до [, разворачивать. Потом собирать число. И добавлять
				 в стек собранную подстроку указанное число раз.
	Tag: Stack
79) 498. Diagonal Traverse
	Линк: https://leetcode.com/problems/diagonal-traverse/
	Сколько раз решал: 1
	Время: 23:59
	Комментарий: Сначала походимся вложенными массивами, кладя кадый элемент в подмассив с индексом диагонали i + j.
				 А потом собираем ответ в цикле, переварачивая нужные диагонали.
	Tag: DiagonalTraverse
80) 438. Find All Anagrams in a String
	Линк: https://leetcode.com/problems/find-all-anagrams-in-a-string/
	Сколько раз решал: 1
	Время: 22:02
	Комментарий: Объединяем использование мапы и окна. В качестве мапы удобнее использовать массив интов, так как не
				 нужно удалять элементы и можно использовать slices.Equals.
	Tag: HashMap, SlidingWindows
81) Implement Heap
	Линк: heap.go, heap_test.go
	Сколько раз решал: 1
	Время: 22:02
	Комментарий: Реализуем методы Len, Less, Swap, Push, Pop.
	Tag: HashMap, SlidingWindows
82) 706. Design HashMap
	Линк: https://leetcode.com/problems/design-hashmap/
	Сколько раз решал: 2
	Время: 18:15
	Комментарий: Реализуем через слайсы, при добавлении проверяем, что текущий размер слайса больше ключа, если нет, то
				 создаем новый слайс разером key + 1 и копируем значаения из предыдущего слайса по условию
				 len(prevSlice) > i, то есть пока тот был заполнен. Оставшиеся элементы заполняем -1. После заполняем
				 полученое занчени.
	Tag: Slice
83) 380. Insert Delete GetRandom O(1)
	Линк: https://leetcode.com/problems/insert-delete-getrandom-o1/
	Сколько раз решал: 1
	Время: 24:40
	Комментарий: Используем мапу и массив. В мапе храним нужное значение и индекс в массиве. В массиве храним приходящие
				 значения. Поэтому тогда, когда нужно получить рандомный, то выбираем рандомный индекс и возвращаем
				 значение из массива. Аккуратно с реализацией удаления! Там нужно перед удалением из мапы переписать
				 удаляемое значение в массиве на последнее, а последнее убрать уменьшением длины массива.
	Tag: HashMap
84) 424. Longest Repeating Character Replacement
	Линк: https://leetcode.com/problems/longest-repeating-character-replacement/
	Сколько раз решал: 1
	Время: 58:20
	Комментарий: Нужна мапа для хранения частоты встречающихся букв. Основная идея в том, что длина окна минус самая
				 большая частота в мапе должна быть меньше, чем k. Если это не так, то двигаем левую границу. Правую же
				 двигаем на каждом шагу и пересчитываем ответ.
	Tag: SlidingWindows
85) 1446. Consecutive Characters
	Линк: https://leetcode.com/problems/consecutive-characters/
	Сколько раз решал: 2
	Время: 04:15
	Комментарий: Ищем наибольшую подстроку с одинаковыми символами.
	Tag: Non-overlapping SlidingWindows
86) 643. Maximum Average Subarray I
	Линк: https://leetcode.com/problems/maximum-average-subarray-i/
	Сколько раз решал: 1
	Время: 08:07
	Комментарий: Ищем максимум в фиксированном окне и возращаем среднее как float64.
	Tag: Fixed SlidingWindows
87) 912. Sort an Array
	Линк: https://leetcode.com/problems/sort-an-array/
	Сколько раз решал: 1
	Время: 36:03
	Комментарий: Сортируем через merge sort. Нужно практиковать.
	Tag: MergeSort
88) 200. Number of Islands
	Линк: https://leetcode.com/problems/number-of-islands/description/
	Сколько раз решал: 1
	Время: 38:48
	Комментарий: Решал через DFS рекурсивно.
	Tag: DFS
89) 62. Unique Paths
	Линк: https://leetcode.com/problems/unique-paths/
	Сколько раз решал: 1
	Время: 39:38
	Комментарий: В каждую клетку можно прийти dp[i][j - 1] + dp[i - 1][j] путями (кладываем результат из левой и верхней)
				 клеток. Если делать по памяти O(m*n), то стоит создать доп строку и столбец с нулями для удобства и
				 не забыть выставить dp[1][1] = 1. Если делать по памяти O(m) или O(n), то стоит помнить, что ответ в
				 tmp[n - 2] или tmp[m - 2], а не в -1, так как первая строка или столбец всегда единицы, то переписав на
				 следующем шаге, мы сдвинем вычисления. Не забудь вначале проверить, что m и n не 1, так как если 1, то
				 счетать не нужно ответ 1 + это сломает решение с O(m) или O(n) по памяти.
	Tag: DP
90) 63. Unique Paths II
	Линк: https://leetcode.com/problems/unique-paths-ii/
	Сколько раз решал: 1
	Время: 34:35
	Комментарий: Решается как и предыдущая, но стоит сбрасывать в 0 значение клетки, если в ней камень, а не считать
				 сумму соседних. Лучше решать через O(m*n) по памяти, так как намного проще.
	Tag: DP
91) 64. Minimum Path Sum
	Линк: https://leetcode.com/problems/minimum-path-sum/
	Сколько раз решал: 1
	Время: 14:44
	Комментарий: Решается как Unique Paths, но только выбираем минимальное значение между левым и верхним, выбранное
				 прибавляем уктекущее значение клетки и записыаем в текущую клетку. Если нельзя изменять входной массив,
				 то стоит создать такой же, только с бортиками math.MaxInt64, что первые сравнения проходили. Не забыть
				 установить в новый массив(если создали) dp[1][1] и не переписывать, так как это начальная точка. Ответ
				 будет в dp[m][n].
	Tag: DP
92) 207. Course Schedule
	Линк: https://leetcode.com/problems/course-schedule/
	Сколько раз решал: 2
	Время: 24:40
	Комментарий: Решал через рекурсивный DFS. Также создал массив смежности, а именно из какого курса переходим
				 в какой. Также создал массив colors, чтобы отмечать цвета: серый = 1, черный = 2. Серый значит, что
				 мы были в этом узле, но можем идти дальше, отмечается в момент прихода в узел. Черный - это значит,
				 что мы обошли все возможные пути для этого узла, отмечается при выходе из узла. Если мы попали в серый
				 узел из серого узла, то значит это цикл и нужно завершаться.
	Tag: DFS
93) 210. Course Schedule II
	Линк: https://leetcode.com/problems/course-schedule-ii/
	Сколько раз решал: 1
	Время: 13:56
	Комментарий: Решается как предыдущая, только после отметки черным узел кладется в результирующий список узлов.
	Tag: DFS
94) 130. Surrounded Regions
	Линк: https://leetcode.com/problems/surrounded-regions/
	Сколько раз решал: 1
	Время: 24:25
	Комментарий: Решал рекурсивным DFS. Сначала запустил DFS для всех краевых узлов, пометив острова у краев посещенными.
				 После запустил DFS, чтобы заметить 'O' на 'X', если это 'O' и если оно не было посещено.
	Tag: DFS
95) 695. Max Area of Island
	Линк: https://leetcode.com/problems/max-area-of-island/
	Сколько раз решал: 1
	Время: 08:03
	Комментарий: Решал рекурсивным DFS. Выбираем максимум из возвращенной dfs функции. Не забываем, перед тем как идти
				 по соседним клеткам в dfs функции пометить, что ты был в текущей клетке, а также создать переменную
				 равную единице - счетчик пройденных подходящих клеток, который потом увеличиваем результатом возвращенным
				 из вызванных dfs функций на соседних клетках.
	Tag: DFS
96) 1091. Shortest Path in Binary Matrix
	Линк: https://leetcode.com/problems/shortest-path-in-binary-matrix/
	Сколько раз решал: 1
	Время: 08:03
	Комментарий: Решал через BFS. В очередь кладем начальную точку (точка - координаты и текущая дистанция). Запускаем
				 цикл, пока очередь не пуста. В цикле достаем первую точку (не забываем убрать ее из очереди queue[1:]
				 (работает даже с одним элементом внутри). Проверяем, что это не финальная точка, если финальная, то 
				 возвращаем дистанцию. Кладем все соседнии точки в очередь, но только если они не выходят валидны (
				 не выходят за границы, не посещались, по ним можно ходить). Помечаем текущую точку посещенной.
				 В конце функции возращаем -1, так как если вышли из цикла, значит не нашли путь до нужной точки.
	Tag: BFS