#include "stdio.h"

int removeElement(int* nums, int numsSize, int val) {
    int j = numsSize - 1;
    int aux = 0;
    int newLength = numsSize;
    for (int i =0; i < numsSize; i++) {
        if (i > j) {
            break;
        }
        if (nums[i] == val) {
            newLength--;
            aux = nums[i];
            nums[i] = nums[j];
            nums[j] = aux;
            j--;
            i--;
        }
    }
    return newLength;
}

int main() {
    printf("Hello, World!\n");
    int nums[] = {3,2,2,3};
    int val = 3;
    int newLength = removeElement(nums, 4, val);
    
    for (int i = 0; i < newLength; i++) {
        printf("%d, ", nums[i]);
    }
    return 0;
}

