// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <iostream>
#include <string>
#include <cstring>
#include <sstream>
#include <vector>

#include "dfl.h"

char* evaluateBool(std::string expresion, int argc, std::string ctx[], int *result) {

  // Convert expression from C++ std::string to C char*
  char *expression_c = new char[expresion.length() + 1];
  std::strcpy(expression_c, expresion.c_str());

  char ** argv = new char *[argc];
  for(unsigned int i = 0; i < argc; i++){
    // Convert each C++ string in context array to C char*
    argv[i] = strdup(ctx[i].c_str());
  }

  // EvaluateBool evaulates the expression for true/false, using the variables provided.
  // The size must be passed with the context array.
  // Returns an error as a char* if any.
  char *err = EvaluateBool(expression_c, argc, argv, result);

  // Free up memory
  free(expression_c);
  delete []argv;

  return err;

}

int main(int argc, char **argv) {

  std::string expression("(@population > 40) and (@featuretype in [road, highway])");
  std::string ctx[] = {"a", "2", "population", "45", "featuretype", "road"};

  // Calculates the size of the ctx array.
  // Since they are char*, we can calculate with the following math.
  int size = sizeof(ctx) / sizeof(ctx[0]);

  int result = 0;

  // Write expresion to stdout
  std::cout << expression << std::endl;

  char *err = evaluateBool(expression, size, ctx, &result);
  if (err != NULL) {
    // Write output to stderr
    std::cerr << std::string(err) << std::endl;
    // Return exit code indicating error
    return 1;
  }

  // Write result to stdout
  std::cout << result << std::endl;

  // Return exit code indicating success
  return 0;
}
