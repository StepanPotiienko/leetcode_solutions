class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        clean_list = sorted(list(set(nums)))
        nums[0:] = clean_list
        return len(clean_list)
