/*
 * @lc app=leetcode id=28 lang=rust
 *
 * [28] Implement strStr()
 */

// @lc code=start
impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        if *needle == "".to_string() {
            return 0;
        }
    
        let hay = haystack.as_bytes();
        let need = needle.as_bytes();
    
        let mut first_char_match_index = -1;
        let mut last_match_index = -1;
    
        for mut i in 0..need.len() {
            let mut j = -1;
            while j < hay.len() as i32 - 1 {
                j += 1;
    
                if i > 0 && last_match_index < 0 {
                    return -1;
                }
    
                if need[i] == hay[j as usize] {
                    if needle.len() == 1 {
                        return j as i32;
                    }
    
                    if i == 0 {
                        first_char_match_index = j as i32;
                        last_match_index = j as i32;
                        i += 1;
                        continue;
                    }
    
                    if last_match_index == j as i32 - 1 {
                        if i == need.len() - 1 {
                            return j - (need.len() - 1) as i32;
                        } else {
                            if j == hay.len() as i32 - 1 {
                                return -1;
                            }
    
                            i += 1;
                            last_match_index = j as i32;
                            continue;
                        }
                    }
                }
    
                if last_match_index >= 0 {
                    i = 0;
                    j = first_char_match_index;
                    last_match_index = -1;
                }
            }
        }
    
        -1
    }
}
// @lc code=end
