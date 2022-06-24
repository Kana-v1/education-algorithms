/*
 * @lc app=leetcode id=104 lang=rust
 *
 * [104] Maximum Depth of Binary Tree
 */

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        pub fn get_node_height(root: Option<&Rc<RefCell<TreeNode>>>) -> i32 {
            if root.is_none() {
                return 0;
            }
        
            let mut left_counter = 0;
            let mut right_counter = 0;
        
            let r = root.unwrap().as_ref().borrow();
        
            if let Some(left_node) = &r.left {
                left_counter += get_node_height(Some(&left_node));
            }
        
            if let Some(right_node) = &r.right {
                right_counter += get_node_height(Some(&right_node));
            }
        
            if left_counter != 0 || right_counter != 0 {
                if left_counter >= right_counter {
                    return left_counter + 1;
                }
        
                return right_counter + 1;
            } else {
                return 1;
            }
        }


        if root.is_none() {
            return  get_node_height(None);
        }
        
        return get_node_height(Some(&root.unwrap()));
    }
}
// @lc code=end

