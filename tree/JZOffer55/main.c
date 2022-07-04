#include<stdio.h>

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
};

int treeHeight(struct TreeNode* root)
{
    int leftH,rightH;

    // 空节点的话，默认为平衡二叉树
    if ( root == NULL )
        return 0;

    // 递归获取左右节点的高度
    leftH = treeHeight(root->left);
    rightH = treeHeight(root->right);

    // 判断左右子节点中是否存在高度大于1的，如果存在则直接返回-1即可，不过不存在判断当前左右节点高度是否大于1
    if ( leftH == -1 || rightH == -1 || fabs(leftH - rightH) > 1)
        return -1;
    else
        return fmax(leftH,rightH) + 1;
}

// 算法主入口
static bool tree_is_balanced(TreeNode* root) {
    // 从低向顶递归的，借鉴与leetcode 的教程
    // 判断 最后的高度是否大于0
    return (treeHeight(root) >= 0);
}

