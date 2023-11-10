#include <stdlib.h>
#include "bubble_sort.c"

#define ARR_LENGTH 1000

int arr1[];

int *genRandomInt32Slice() {
    int *res = malloc(ARR_LENGTH * sizeof(int));

    for (int i = 0; i < ARR_LENGTH; i++) {
        *(res+i) = rand();
    }

    return res;
}
//func genRandomInt32Slice(res *[]int32) {
//    for i := 0; i < len(*res); i++ {
//        (*res)[i] = rand.Int31n(math.MaxInt32)
//    }
//}


int main() {
    int *a = genRandomInt32Slice();
    bubbleSortDesc(a, sizeof(a) / sizeof(a[0]));
    for (int i = 0; i < sizeof(a) / sizeof(a[0]); i++) {
        printf("%d\n", a[i]);
    }
}