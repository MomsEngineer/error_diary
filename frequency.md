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

21) 304. Range Sum Query 2D - Immutable
22) 17. Letter Combinations of a Phone Number
23) 22. Generate Parentheses
24) 977. Squares of a Sorted Array
25) 125. Valid Palindrome
26) 167. Two Sum II - Input Array Is Sorted
27) 15. 3Sum
28) 16. 3Sum Closest
29) 11. Container With Most Water
30) Intersection Of Sorted Arrays (https://www.interviewbit.com/problems/intersection-of-sorted-arrays/)
31) 844. Backspace String Compare
32) 283. Move Zeroes
33) 14. Longest Common Prefix
34) One Edit Distance
35) 228. Summary Ranges
36) 485. Max Consecutive Ones
37) 849. Maximize Distance to Closest Person
38) 443. String Compression
39) 1004. Max Consecutive Ones III
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
	Сколько раз решал: 1
	Время: 05:02
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


**Классические обходы деревьев**

83. Remove Duplicates from Sorted List
23. Merge k Sorted Lists

- [ ]  [144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) - [РЕШЕНИЕ](https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/989075817/)
- [ ]  [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)
- [ ]  [145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)

**Обход по уровням**

- [ ]  [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) - [РЕШЕНИЕ](https://leetcode.com/problems/binary-tree-level-order-traversal/submissions/989091886/)
- [ ]  [**103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/) #ВАЖНО**
- [ ]  [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/) - [РЕШЕНИЕ](https://leetcode.com/problems/binary-tree-right-side-view/submissions/989084003/)

**Деревья**

- [ ]  [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) - [РЕШЕНИЕ](https://leetcode.com/problems/symmetric-tree/submissions/969526824/) **#ВАЖНО**
- [ ]  [100. Same Tree](https://leetcode.com/problems/same-tree/)  - [РЕШЕНИЕ](https://leetcode.com/problems/same-tree/submissions/1025715566/)

**Сумма вершин дерева**

- [ ]  [112. Path Sum](https://leetcode.com/problems/path-sum/) - [РЕШЕНИЕ](https://leetcode.com/problems/path-sum/submissions/989133327/) **#ВАЖНО**

Доп
498. Diagonal Traverse
438. Find All Anagrams in a String
706. Design HashMap
380. Insert Delete GetRandom O(1)
146. LRU Cache
460. LFU Cache
424. Longest Repeating Character Replacement
1446. Consecutive Characters
643. Maximum Average Subarray I
Implement Heap
912. Sort an Array