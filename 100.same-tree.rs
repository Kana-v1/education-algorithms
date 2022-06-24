/*
 * @lc app=leetcode id=100 lang=rust
 *
 * [100] Same Tree
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
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn is_same_tree(
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        pub fn graph_triversal(g: Option<Rc<RefCell<TreeNode>>>) -> Vec<Option<i32>> {
            let mut buf = Vec::<Option<i32>>::new();

            if g.is_none() {
                return vec![None]
            }
    
            let val = get_node_values(Some(&g.unwrap()), &mut buf);
    
            let mut res = vec![val];
    
            res.extend(buf);
    
            res
        }
    
        pub fn get_node_values(
            g: Option<&Rc<RefCell<TreeNode>>>,
            stack: &mut Vec<Option<i32>>,
        ) -> Option<i32> {
            match g {
                None => {
                    panic!("get null tree node")
                }
                Some(gr) => {
                    let graph = gr.borrow();
    
                    match &graph.left {
                        Some(left_node) => {
                            let val = get_node_values(Some(left_node), stack).clone();
                            stack.push(val);
                        }
    
                        None => {
                            if graph.right.is_some() {
                                stack.push(None);
                            }
                        }
                    }
    
                    if let Some(right_node) = &graph.right {
                        let val = get_node_values(Some(right_node), stack).clone();
                        stack.push(val);
                    }
    
                    Some(gr.borrow().val.clone())
                }
            }
        }
    
        let p_els = graph_triversal(p);
        let q_els = graph_triversal(q);
    
        if p_els.len() != q_els.len() {
            return false;
        }
    
        for i in 0..p_els.len() {
            if p_els[i].is_none() && !q_els[i].is_none() || q_els[i].is_none() && !p_els[i].is_none() {
                return false;
            }
    
            if p_els[i].is_none() && q_els[i].is_none() {
                continue;
            }

            if p_els[i].unwrap() != q_els[i].unwrap() {
                return false;
            }
        }
    
        true
    }
}
// @lc code=end
