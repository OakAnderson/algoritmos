Golang Data Structures and Algorithms
=========================================

Exemplos simples de implementação de estrutura de dados e algoritmos em Go.

Você pode testar criando um arquivo go: (Ex: uso de `BubbleSort` em `sort`)

```go
package main

import (
    "gitlab.com/OakAnderson/algoritmos/algoritmos/sort"
)

func main() {
    myList := []int{1, 8, 3, 5, 6}
    myList = sort.BubbleSort(myList)
    print(myList)
}
```


## List of Implementations

- [~~arrays~~](algorithms/arrays)
<!--    - [~~delete_nth~~](algorithms/arrays/delete_nth.go)
    - [~~flatten~~](algorithms/arrays/flatten.go)
    - [~~garage~~](algorithms/arrays/garage.go)
    - [~~josephus_problem~~](algorithms/arrays/josephus.go)
    - [~~limit~~](algorithms/arrays/limit.go)
    - [~~longest_non_repeat~~](algorithms/arrays/longest_non_repeat.go/)
    - [~~max_ones_index~~](algorithms/arrays/max_ones_index.go)
    - [~~merge_intervals~~](algorithms/arrays/merge_intervals.go)
    - [~~missing_ranges~~](algorithms/arrays/missing_ranges.go)
    - [~~plus_one~~](algorithms/arrays/plus_one.go)
    - [~~rotate~~](algorithms/arrays/rotate.go)
    - [~~summarize_ranges~~](algorithms/arrays/summarize_ranges.go)
    - [~~three_sum~~](algorithms/arrays/three_sum.go)
    - [~~trimmean~~](algorithms/arrays/trimmean.go)
    - [~~top_1~~](algorithms/arrays/top_1.go)
    - [~~two_sum~~](algorithms/arrays/two_sum.go)
    - [~~move_zeros~~](algorithms/arrays/move_zeros.go)
    - [~~n_sum~~](algorithms/arrays/n_sum.go) -->
- [~~automata~~](algorithms/automata)
<!--    - [~~DFA~~](algorithms/automata/DFA.go) -->
- [~~backtrack~~](algorithms/backtrack)
<!--    - [~~general_solution~~.md](algorithms/backtrack/)
    - [~~add_operators~~](algorithms/backtrack/add_operators.go)
    - [~~anagram~~](algorithms/backtrack/anagram.go)
    - [~~array_sum_combinations~~](algorithms/backtrack/array_sum_combinations.go)
    - [~~combination_sum~~](algorithms/backtrack/combination_sum.go)
    - [~~factor_combinations~~](algorithms/backtrack/factor_combinations.go)
    - [~~generate_abbreviations~~](algorithms/backtrack/generate_abbreviations.go)
    - [~~generate_parenthesis~~](algorithms/backtrack/generate_parenthesis.go)
    - [~~letter_combination~~](algorithms/backtrack/letter_combination.go)
    - [~~palindrome_partitioning~~](algorithms/backtrack/palindrome_partitioning.go)
    - [~~pattern_match~~](algorithms/backtrack/pattern_match.go)
    - [~~permute~~](algorithms/backtrack/permute.go)
    - [~~permute_unique~~](algorithms/backtrack/permute_unique.go)
    - [~~subsets~~](algorithms/backtrack/subsets.go)
    - [~~subsets_unique~~](algorithms/backtrack/subsets_unique.go) -->
- [~~bfs~~](algorithms/bfs)
    <!-- - [~~maze_search~~](algorithms/bfs/maze_search.go)
    - [~~shortest_distance_from_all_buildings~~](algorithms/bfs/shortest_distance_from_all_buildings.go)
    - [~~word_ladder~~](algorithms/bfs/word_ladder.go) -->
- [~~bit~~](algorithms/bit)
    <!-- - [~~add_bitwise_operator~~](algorithms/bit/add_bitwise_operator.go)
    - [~~bit_operation~~](algorithms/bit/bit_operation.go)
    - [~~bytes_int_conversion~~](algorithms/bit/bytes_int_conversion.go)
    - [~~count_flips_to_convert~~](algorithms/bit/count_flips_to_convert.go)
    - [~~count_ones~~](algorithms/bit/count_ones.go)
    - [~~find_difference~~](algorithms/bit/find_difference.go)
    - [~~find_missing_number~~](algorithms/bit/find_missing_number.go)
    - [~~flip_bit_longest_sequence~~](algorithms/bit/flip_bit_longest_sequence.go)
    - [~~power_of_two~~](algorithms/bit/power_of_two.go)
    - [~~reverse_bits~~](algorithms/bit/reverse_bits.go)
    - [~~single_number~~](algorithms/bit/single_number.go)
    - [~~single_number2~~](algorithms/bit/single_number2.go)
    - [~~single_number3~~](algorithms/bit/single_number3.go)
    - [~~subsets~~](algorithms/bit/subsets.go)
    - [~~swap_pair~~](algorithms/bit/swap_pair.go)
    - [~~has_alternative_bit~~](algorithms/bit/has_alternative_bit.go)
    - [~~insert_bit~~](algorithms/bit/insert_bit.go)
    - [~~remove_bit~~](algorithms/bit/remove_bit.go)
    - [~~binary_gap~~](algorithms/bit/binary_gap.go) -->
- [~~compression~~](algorithms/compression)
    <!-- - [~~huffman_coding~~](algorithms/compression/huffman_coding.go)
    - [~~rle_compression~~](algorithms/compression/rle_compression.go)
    - [~~elias~~](algorithms/compression/elias.go) -->
- [~~dfs~~](algorithms/dfs)
    <!-- - [~~all_factors~~](algorithms/dfs/all_factors.go)
    - [~~count_islands~~](algorithms/dfs/count_islands.go)
    - [~~pacific_atlantic~~](algorithms/dfs/pacific_atlantic.go)
    - [~~sudoku_solver~~](algorithms/dfs/sudoku_solver.go)
    - [~~walls_and_gates~~](algorithms/dfs/walls_and_gates.go) -->
- [~~distribution~~](algorithms/distribution)
    <!-- - [~~histogram~~](algorithms/distribution/histogram.go) -->
- [~~dp~~](algorithms/dp)
    <!-- - [~~buy_sell_stock~~](algorithms/dp/buy_sell_stock.go)
    - [~~climbing_stairs~~](algorithms/dp/climbing_stairs.go)
    - [~~coin_change~~](algorithms/dp/coin_change.go)
    - [~~combination_sum~~](algorithms/dp/combination_sum.go)
    - [~~egg_drop~~](algorithms/dp/egg_drop.go)
    - [~~house_robber~~](algorithms/dp/house_robber.go)
    - [~~int_divide~~](algorithms/dp/int_divide.go)
    - [~~job_scheduling~~](algorithms/dp/job_scheduling.go)
    - [~~knapsack~~](algorithms/dp/knapsack.go)
    - [~~longest_increasing~~](algorithms/dp/longest_increasing.go)
    - [~~matrix_chain_order~~](algorithms/dp/matrix_chain_order.go)
    - [~~max_product_subarray~~](algorithms/dp/max_product_subarray.go)
    - [~~max_subarray~~](algorithms/dp/max_subarray.go)
    - [~~min_cost_path~~](algorithms/dp/min_cost_path.go)
    - [~~num_decodings~~](algorithms/dp/num_decodings.go)
    - [~~regex_matching~~](algorithms/dp/regex_matching.go)
    - [~~rod_cut~~](algorithms/dp/rod_cut.go)
    - [~~word_break~~](algorithms/dp/word_break.go)
    - [~~fibonacci~~](algorithms/dp/fib.go)
	- [~~hosoya~~ triangle](algorithms/dp/hosoya_triangle.go) -->
- [~~graph~~](algorithms/graph)
    <!-- - [~~check_bipartite~~](algorithms/graph/check_bipartite.go)
    - [~~strongly_connected~~](algorithms/graph/check_digraph_strongly_connected.go)
    - [~~clone_graph~~](algorithms/graph/clone_graph.go)
    - [~~cycle_detection~~](algorithms/graph/cycle_detection.go)
    - [~~find_all_cliques~~](algorithms/graph/find_all_cliques.go)
    - [~~find_path~~](algorithms/graph/find_path.go)
    - [~~graph~~](algorithms/graph/graph.go)
    - [~~dijkstra~~](algorithms/graph/dijkstra.go)
    - [~~markov_chain~~](algorithms/graph/markov_chain.go)
    - [~~minimum_spanning_tree~~](algorithms/graph/minimum_spanning_tree.go)
    - [~~satisfiability~~](algorithms/graph/satisfiability.go)
    - [~~tarjan~~](algorithms/graph/tarjan.go)
    - [~~traversal~~](algorithms/graph/traversal.go)
	  - [~~maximum_flow~~](algorithms/graph/maximum_flow.go)
    - [~~maximum_flow_bfs~~](algorithms/graph/maximum_flow_bfs.go)
    - [~~maximum_flow_dfs~~](algorithms/graph/maximum_flow_dfs.go)
    - [~~all_pairs_shortest_path~~](algorithms/graph/all_pairs_shortest_path.go)
    - [~~bellman_ford~~](algorithms/graph/bellman_ford.go)
    - [~~Count~~ Connected Components](algoritms/graph/count_connected_number_of_component.go) -->
- [~~heap~~](algorithms/heap)
    <!-- - [~~merge_sorted_k_lists~~](algorithms/heap/merge_sorted_k_lists.go)
    - [~~skyline~~](algorithms/heap/skyline.go)
    - [~~sliding_window_max~~](algorithms/heap/sliding_window_max.go)
    - [~~binary_heap~~](algorithms/heap/binary_heap.go)
    - [~~k_closest_points~~](algorithms/heap/k_closest_points.go) -->
- [~~linkedlist~~](algorithms/linkedlist)
    <!-- - [~~add_two_numbers~~](algorithms/linkedlist/add_two_numbers.go)
    - [~~copy_random_pointer~~](algorithms/linkedlist/copy_random_pointer.go)
    - [~~delete_node~~](algorithms/linkedlist/delete_node.go)
    - [~~first_cyclic_node~~](algorithms/linkedlist/first_cyclic_node.go)
    - [~~is_cyclic~~](algorithms/linkedlist/is_cyclic.go)
    - [~~is_palindrome~~](algorithms/linkedlist/is_palindrome.go)
    - [~~kth_to_last~~](algorithms/linkedlist/kth_to_last.go)
    - [~~linkedlist~~](algorithms/linkedlist/linkedlist.go)
    - [~~remove_duplicates~~](algorithms/linkedlist/remove_duplicates.go)
    - [~~reverse~~](algorithms/linkedlist/reverse.go)
    - [~~rotate_list~~](algorithms/linkedlist/rotate_list.go)
    - [~~swap_in_pairs~~](algorithms/linkedlist/swap_in_pairs.go)
    - [~~is_sorted~~](algorithms/linkedlist/is_sorted.go)
    - [~~remove_range~~](algorithms/linkedlist/remove_range.go) -->
- [~~map~~](algorithms/map)
    <!-- - [~~hashtable~~](algorithms/map/hashtable.go)
    - [~~separate_chaining_hashtable~~](algorithms/map/separate_chaining_hashtable.go)
    - [~~longest_common_subsequence~~](algorithms/map/longest_common_subsequence.go)
    - [~~randomized_set~~](algorithms/map/randomized_set.go)
    - [~~valid_sudoku~~](algorithms/map/valid_sudoku.go)
    - [~~word_pattern~~](algorithms/map/word_pattern.go)
    - [~~is_isomorphic~~](algorithms/map/is_isomorphic.go)
    - [~~is_anagram~~](algorithms/map/is_anagram.go)     -->
- [~~maths~~](algorithms/maths)
    <!-- - [~~base_conversion~~](algorithms/maths/base_conversion.go)
    - [~~combination~~](algorithms/maths/combination.go)
    - [~~cosine_similarity~~](algorithms/maths/cosine_similarity.go)
    - [~~decimal_to_binary_ip~~](algorithms/maths/decimal_to_binary_ip.go)
    - [~~euler_totient~~](algorithms/maths/euler_totient.go)
    - [~~extended_gcd~~](algorithms/maths/extended_gcd.go)
    - [~~factorial~~](algorithms/maths/factorial.go)    
    - [~~gcd~~/lcm](algorithms/maths/gcd.go)
    - [~~generate_strobogrammtic~~](algorithms/maths/generate_strobogrammtic.go)
    - [~~is_strobogrammatic~~](algorithms/maths/is_strobogrammatic.go)
    - [~~modular_exponential~~](algorithms/maths/modular_exponential.go)
    - [~~next_bigger~~](algorithms/maths/next_bigger.go)
    - [~~next_perfect_square~~](algorithms/maths/next_perfect_square.go)
    - [~~nth_digit~~](algorithms/maths/nth_digit.go)
    - [~~prime_check~~](algorithms/maths/prime_check.go)
    - [~~primes_sieve_of_eratosthenes~~](algorithms/maths/primes_sieve_of_eratosthenes.go)
    - [~~pythagoras~~](algorithms/maths/pythagoras.go)
    - [~~rabin_miller~~](algorithms/maths/rabin_miller.go)
    - [~~rsa~~](algorithms/maths/rsa.go)
    - [~~sqrt_precision_factor~~](algorithms/maths/sqrt_precision_factor.go)
    - [~~summing_digits~~](algorithms/maths/summing_digits.go)
    - [~~hailstone~~](algorithms/maths/hailstone.go)
    - [~~recursive_binomial_coefficient~~](algorithms/maths/recursive_binomial_coefficient.go)
    - [~~find_order~~](algorithms/maths/find_order_simple.go)
	  - [~~find_primitive_root~~](algorithms/maths/find_primitive_root_simple.go)
	  - [~~diffie_hellman_key_exchange~~](algorithms/maths/diffie_hellman_key_exchange.go) -->
- [~~matrix~~](algorithms/matrix)
    <!-- - [~~sudoku_validator~~](algorithms/matrix/sudoku_validator.go)
    - [~~bomb_enemy~~](algorithms/matrix/bomb_enemy.go)
    - [~~copy_transform~~](algorithms/matrix/copy_transform.go)
    - [~~count_paths~~](algorithms/matrix/count_paths.go)
    - [~~matrix_rotation~~.txt](algorithms/matrix/matrix_rotation.txt)
    - [~~matrix_inversion~~](algorithms/matrix/matrix_inversion.go)
    - [~~matrix_multiplication~~](algorithms/matrix/multiply.go)
    - [~~rotate_image~~](algorithms/matrix/rotate_image.go)
    - [~~search_in_sorted_matrix~~](algorithms/matrix/search_in_sorted_matrix.go)
    - [~~sparse_dot_vector~~](algorithms/matrix/sparse_dot_vector.go)
    - [~~sparse_mul~~](algorithms/matrix/sparse_mul.go)
    - [~~spiral_traversal~~](algorithms/matrix/spiral_traversal.go)
	- [~~crout_matrix_decomposition~~](algorithms/matrix/crout_matrix_decomposition.go)
	- [~~cholesky_matrix_decomposition~~](algorithms/matrix/cholesky_matrix_decomposition.go)
    - [~~sum_sub_squares~~](algorithms/matrix/sum_sub_squares.go) -->
- [~~queues~~](algorithms/queues)
    <!-- - [~~max_sliding_window~~](algorithms/queues/max_sliding_window.go)
    - [~~moving_average~~](algorithms/queues/moving_average.go)
    - [~~queue~~](algorithms/queues/queue.go)
    - [~~reconstruct_queue~~](algorithms/queues/reconstruct_queue.go)
    - [~~zigzagiterator~~](algorithms/queues/zigzagiterator.go) -->
- [~~search~~](algorithms/search)
    <!-- - [~~binary_search~~](algorithms/search/binary_search.go)
    - [~~first_occurrence~~](algorithms/search/first_occurrence.go)
    - [~~last_occurrence~~](algorithms/search/last_occurrence.go)
    - [~~linear_search~~](algorithms/search/linear_search.go)
    - [~~search_insert~~](algorithms/search/search_insert.go)
    - [~~two_sum~~](algorithms/search/two_sum.go)
    - [~~search_range~~](algorithms/search/search_range.go)
    - [~~find_min_rotate~~](algorithms/search/find_min_rotate.go)
    - [~~search_rotate~~](algorithms/search/search_rotate.go)
    - [~~jump_search~~](algorithms/search/jump_search.go)
    - [~~next_greatest_letter~~](algorithms/search/next_greatest_letter.go) -->
- [~~set~~](algorithms/set)
    <!-- - [~~randomized_set~~](algorithms/set/randomized_set.go)
    - [~~set_covering~~](algorithms/set/set_covering.go)
    - [~~find_keyboard_row~~](algorithms/set/find_keyboard_row.go) -->
- [sort](algorithms/sort)
    - [BitonicSort](algorithms/sort/bitonicSort.go)
    - [BitonicSortMultiprocess](algorithms/sort/bitonicSortMultiprocess.go)
    - [BubbleSort](algorithms/sort/bubbleSort.go)
    - [InsertionSort](algorithms/sort/insertionSort.go)
    - [~~bogoSort~~](algorithms/sort/bogoSort.go)
    - [~~bucketSort~~](algorithms/sort/bucketSort.go)
    - [~~cocktailShakerSort~~](algorithms/sort/cocktailShakerSort.go)
    - [~~combSort~~](algorithms/sort/combSort.go)
    - [~~countingSort~~](algorithms/sort/countingSort.go)
    - [~~cycleSort~~](algorithms/sort/cycleSort.go)
    - [~~gnomeSort~~](algorithms/sort/gnomeSort.go)
    - [~~heapSort~~](algorithms/sort/heapSort.go)
    - [~~meetingRooms~~](algorithms/sort/meetingRooms.go)
    - [~~mergeSort~~](algorithms/sort/mergeSort.go)
    - [~~pancakeSort~~](algorithms/sort/pancakeSort.go)
    - [~~quickSort~~](algorithms/sort/quickSort.go)
    - [~~radixSort~~](algorithms/sort/radixSort.go)
    - [~~selectionSort~~](algorithms/sort/selectionSort.go)
    - [~~shellSort~~](algorithms/sort/shellSort.go)
    - [~~sortColors~~](algorithms/sort/sortColors.go)
    - [~~topSort~~](algorithms/sort/topSort.go)
    - [~~wiggleSort~~](algorithms/sort/wiggleSort.go)
- [~~stack~~](algorithms/stack)
    <!-- - [~~longest_abs_path~~](algorithms/stack/longest_abs_path.go)
    - [~~simplify_path~~](algorithms/stack/simplify_path.go)
    - [~~stack~~](algorithms/stack/stack.go)
    - [~~valid_parenthesis~~](algorithms/stack/valid_parenthesis.go)
    - [~~stutter~~](algorithms/stack/stutter.go)
    - [~~switch_pairs~~](algorithms/stack/switch_pairs.go)
    - [~~is_consecutive~~](algorithms/stack/is_consecutive.go)
    - [~~remove_min~~](algorithms/stack/remove_min.go)
    - [~~is_sorted~~](algorithms/stack/is_sorted.go) -->
- [~~strings~~](algorithms/strings)
    <!-- - [~~fizzbuzz~~](algorithms/strings/fizzbuzz.go)
    - [~~delete_reoccurring~~](algorithms/strings/delete_reoccurring.go)
    - [~~strip_url_params~~](algorithms/strings/strip_url_params.go)
    - [~~validate_coordinates~~](algorithms/strings/validate_coordinates.go)
    - [~~domain_extractor~~](algorithms/strings/domain_extractor.go)
    - [~~merge_string_checker~~](algorithms/strings/merge_string_checker.go)
    - [~~add_binary~~](algorithms/strings/add_binary.go)
    - [~~breaking_bad~~](algorithms/strings/breaking_bad.go)
    - [~~decode_string~~](algorithms/strings/decode_string.go)
    - [~~encode_decode~~](algorithms/strings/encode_decode.go)
    - [~~group_anagrams~~](algorithms/strings/group_anagrams.go)
    - [~~int_to_roman~~](algorithms/strings/int_to_roman.go)
    - [~~is_palindrome~~](algorithms/strings/is_palindrome.go)
    - [~~license_number~~](algorithms/strings/license_number.go)
    - [~~make_sentence~~](algorithms/strings/make_sentence.go)
    - [~~multiply_strings~~](algorithms/strings/multiply_strings.go)
    - [~~one_edit_distance~~](algorithms/strings/one_edit_distance.go)
    - [~~rabin_karp~~](algorithms/strings/rabin_karp.go)
    - [~~reverse_string~~](algorithms/strings/reverse_string.go)
    - [~~reverse_vowel~~](algorithms/strings/reverse_vowel.go)
    - [~~reverse_words~~](algorithms/strings/reverse_words.go)
    - [~~roman_to_int~~](algorithms/strings/roman_to_int.go)
    - [~~word_squares~~](algorithms/strings/word_squares.go)
    - [~~unique_morse~~](algorithms/strings/unique_morse.go)
    - [~~judge_circle~~](algorithms/strings/judge_circle.go)
    - [~~strong_password~~](algorithms/strings/strong_password.go)
    - [~~caesar_cipher~~](algorithms/strings/caesar_cipher.go)
    - [~~contain_string~~](algorithms/strings/contain_string.go)
    - [~~count_binary_substring~~](algorithms/strings/count_binary_substring.go)
    - [~~repeat_string~~](algorithms/strings/repeat_string.go)
    - [~~min_distance~~](algorithms/strings/min_distance.go)
    - [~~longest_common_prefix~~](algorithms/strings/longest_common_prefix.go)
    - [~~rotate~~](algorithms/strings/rotate.go)
    - [~~first_unique_char~~](algorithms/strings/first_unique_char.go)
    - [~~repeat_substring~~](algorithms/strings/repeat_substring.go)     
	- [~~atbash_cipher~~](algorithms/strings/atbash_cipher.go)
	- [~~knuth_morris_pratt~~](algorithms/strings/knuth_morris_pratt.go) -->
- [~~tree~~](algorithms/tree)
    <!-- - [~~bst~~](algorithms/tree/bst)
        - [~~array_to_bst~~](algorithms/tree/bst/array_to_bst.go)
        - [~~bst_closest_value~~](algorithms/tree/bst/bst_closest_value.go)
        - [~~BSTIterator~~](algorithms/tree/bst/BSTIterator.go)
        - [~~delete_node~~](algorithms/tree/bst/delete_node.go)
        - [~~is_bst~~](algorithms/tree/bst/is_bst.go)
        - [~~kth_smallest~~](algorithms/tree/bst/kth_smallest.go)
        - [~~lowest_common_ancestor~~](algorithms/tree/bst/lowest_common_ancestor.go)
        - [~~predecessor~~](algorithms/tree/bst/predecessor.go)
        - [~~serialize_deserialize~~](algorithms/tree/bst/serialize_deserialize.go)
        - [~~successor~~](algorithms/tree/bst/successor.go)
        - [~~unique_bst~~](algorithms/tree/bst/unique_bst.go)
        - [~~depth_sum~~](algorithms/tree/bst/depth_sum.go)
        - [~~count_left_node~~](algorithms/tree/bst/count_left_node.go)
        - [~~num_empty~~](algorithms/tree/bst/num_empty.go)
        - [~~height~~](algorithms/tree/bst/height.go)
    - [~~red_black_tree~~](algorithms/tree/red_black_tree)
        - [~~red_black_tree~~](algorithms/tree/red_black_tree/red_black_tree.go)
    - [~~segment_tree~~](algorithms/tree/segment_tree)
        - [~~segment_tree~~](algorithms/tree/segment_tree/segment_tree.go)
        - [~~iterative_segment_tree~~](algorithms/tree/segment_tree/iterative_segment_tree.go)
    - [~~traversal~~](algorithms/tree/traversal)
        - [~~inorder~~](algorithms/tree/traversal/inorder.go)
        - [~~level_order~~](algorithms/tree/traversal/level_order.go)
        - [~~postorder~~](algorithms/tree/traversal/postorder.go)
        - [~~preorder~~](algorithms/tree/traversal/preorder.go)
        - [~~zigzag~~](algorithms/tree/traversal/zigzag.go)
    - [~~trie~~](algorithms/tree/trie)
        - [~~add_and_search~~](algorithms/tree/trie/add_and_search.go)
        - [~~trie~~](algorithms/tree/trie/trie.go)
    - [~~b_tree~~](algorithms/tree/b_tree.go)
    - [~~binary_tree_paths~~](algorithms/tree/binary_tree_paths.go)
    - [~~bin_tree_to_list~~](algorithms/tree/bin_tree_to_list.go)
    - [~~deepest_left~~](algorithms/tree/deepest_left.go)
    - [~~invert_tree~~](algorithms/tree/invert_tree.go)
    - [~~is_balanced~~](algorithms/tree/is_balanced.go)
    - [~~is_subtree~~](algorithms/tree/is_subtree.go)
    - [~~is_symmetric~~](algorithms/tree/is_symmetric.go)
    - [~~longest_consecutive~~](algorithms/tree/longest_consecutive.go)
    - [~~lowest_common_ancestor~~](algorithms/tree/lowest_common_ancestor.go)
    - [~~max_height~~](algorithms/tree/max_height.go)
    - [~~max_path_sum~~](algorithms/tree/max_path_sum.go)
    - [~~min_height~~](algorithms/tree/min_height.go)
    - [~~path_sum~~](algorithms/tree/path_sum.go)
    - [~~path_sum2~~](algorithms/tree/path_sum2.go)
    - [~~pretty_print~~](algorithms/tree/pretty_print.go)
    - [~~same_tree~~](algorithms/tree/same_tree.go)
    - [~~tree~~](algorithms/tree/tree.go) -->
- [~~unix~~](algorithms/unix)
    <!-- - [~~path~~](algorithms/unix/path/)
        - [~~join_with_slash~~](algorithms/unix/path/join_with_slash.go)
        - [~~full_path~~](algorithms/unix/path/full_path.go)
        - [~~split~~](algorithms/unix/path/split.go)
        - [~~simplify_path~~](algorithms/unix/path/simplify_path.go) -->
- [~~unionfind~~](algorithms/unionfind)
    <!-- - [~~count_islands~~](algorithms/unionfind/count_islands.go) -->
