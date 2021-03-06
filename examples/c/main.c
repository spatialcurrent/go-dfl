// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "dfl.h"

int
main(int argc, char **argv) {
    char *err;

    //
    char *expression = "(@population > 40) and (@featuretype in [road, highway])";

    // You must prepare your context as an alternating list of keys and values.
    /// TryConvertString attempts to convert any values into their appropriate type.
    // For example the below values get converted into:
    // {"a":2, "population": 45, "featuretype": "road"}
    char size = 6;
    char *ctx[] = {"a", "2", "population", "45", "featuretype", "road"};

    printf("ctx %s\n", ctx[2]);

    char** args;
    args = malloc(size * sizeof(char*));
    for(unsigned int i = 0; i < size; i++){
        args[i] = ctx[i];
    }

    // Declare a variable to store the result.
    // 0 = false and 1 = true
    int result = 0;

    printf("%s\n", expression);

    // EvaluateBool evaulates the expression for true/false, using the variables provided.
    // The size must be passed with the context array.
    // Returns an error as a string if any.
    err = EvaluateBool(expression, size, args, &result);
    if (err != NULL) {
        fprintf(stderr, "error: %s\n", err);
        free(err);
        return 1;
    }

    printf("%d\n", result);

    return 0;
}
