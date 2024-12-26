1) 1. Two Sum
	Линк: https://leetcode.com/problems/two-sum/
	Сколько раз решал: 1
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
	Сколько раз решал: 1
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
	Сколько раз решал: 1
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
	Комментарий: XOR (go ^), обратная операция XOR - это XOR
	Tag: PrefixArray

?????????????????????????????
21) 304. Range Sum Query 2D - Immutable

22) 17. Letter Combinations of a Phone Number
	Линк: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
	Сколько раз решал: 1
	Время: 29:37
	Комментарий: Полный обход в рекурсии.
	Tag: Backtracking

?????????????????????????????
23) 22. Generate Parentheses

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
34) One Edit Distance (https://www.lintcode.com/problem/640/)

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
41) 567. Permutation in String

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

?????????????????????????????
50) 81. Search in Rotated Sorted Array II
	Линк: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
	Сколько раз решал:
	Время:
	Комментарий:
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
	Сколько раз решал: 1
	Время: 14:42
	Комментарий: Используй Dummy, чтобы не мучится с проверками на коротких списках.
	Tag: LinkedList
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
	Сколько раз решал: 1
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

67) 20. Valid Parentheses
	Линк: https://leetcode.com/problems/valid-parentheses/
	Сколько раз решал: 1
	Время: 12:05
	Комментарий: Пиши внимательно и не торопясь. Обрати внимание на последнюю проверку, в мапе с парами ищем скобку из
				 стека, а не из строки, так как из строки закрывающая.
	Tag: Stack
68) 20. Valid Parentheses
	Линк: https://leetcode.com/problems/daily-temperatures/
	Сколько раз решал: 1
	Время: 19:46
	Комментарий: Много логических ошибок. Решай внимательней и проще.
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
79) 498. Diagonal Traverse
80) 438. Find All Anagrams in a String
81) Implement Heap
82) 706. Design HashMap
83) 380. Insert Delete GetRandom O(1)
84) 146. LRU Cache
85) 460. LFU Cache
86) 424. Longest Repeating Character Replacement
87) 1446. Consecutive Characters
88) 643. Maximum Average Subarray I
89) 912. Sort an Array