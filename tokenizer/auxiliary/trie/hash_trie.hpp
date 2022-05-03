#ifndef HASH_TRIE_HPP
#define HASH_TRIE_HPP

#include <vector>
#include <set>
#include <string>
#include "../utf8.h"
#include "../tsl/robin_set.h"
#include <mutex>

template < class Node >
struct HashTrie
{
	template < typename T >
	using fast_hash_set_t = tsl::robin_set< T >;

	std::vector< Node > pool;
	fast_hash_set_t< uint32_t > alphabet;
	std::shared_mutex mtwrite;

	HashTrie()
	{
		pool.assign(1, Node()); // root
	}

	inline void add_child(int u, uint32_t c)
	{
		pool[u].children[c] = pool.size();
		pool.push_back(Node());

		std::lock_guard<std::shared_mutex> lock(mtwrite);
		alphabet.insert(c);
	}

	inline bool has_child(int u, uint32_t c)
	{
		return pool[u].children.find(c) != pool[u].children.end();
	}

	inline int get_child(int u, uint32_t c)
	{
		return pool[u].children[c];
	}

	int add_new_term(const std::string &s, int frequency)
	{
		int cur_node = 0;
		utf8::unchecked::iterator< std::string::const_iterator > it(s.begin()), end_it(s.end());
		while (it != end_it)
		{
			uint32_t codepoint = *it;
			if (!has_child(cur_node, codepoint))
			{
				add_child(cur_node, codepoint);
			}
			cur_node = get_child(cur_node, codepoint);
			it++;
		}
		pool[cur_node].mark_ending(frequency);
		return cur_node;
	}

	std::vector< Node > &dump_structure()
	{
		for (Node &u : pool)
		{
			if (~u.frequency)
			{
				u.finalize();
			}
		}
		return pool;
	}

	std::set< uint32_t > dump_alphabet()
	{
		std::shared_lock<std::shared_mutex> lock(mtwrite);
		return std::set< uint32_t >(alphabet.begin(), alphabet.end());
	}
};

#endif // HASH_TRIE_HPP