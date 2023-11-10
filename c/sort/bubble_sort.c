#include <stdio.h>

void bubbleSortAsc(int *arr, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (*(arr + i) > *(arr + j)) {
               int tmp = *(arr + j);
               *(arr+j) = *(arr + i);
                *(arr + i) = tmp;
            }
        }
    }
}

void bubbleSortDesc(int *arr, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = i + 1; j < length; j++) {
            if (*(arr + i) < *(arr + j)) {
                int tmp = *(arr + j);
                *(arr+j) = *(arr + i);
                *(arr + i) = tmp;
            }
        }
    }
}
