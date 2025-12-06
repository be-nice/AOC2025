#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static char *read_file(const char *filename) {
  FILE *fp = fopen(filename, "r");
  if (!fp) {
    perror("File open error");
    exit(1);
  }

  fseek(fp, 0, SEEK_END);
  long size = ftell(fp);
  fseek(fp, 0, SEEK_SET);

  char *buffer = malloc((size_t)size + 1);
  if (!buffer) {
    perror("malloc");
    exit(1);
  }

  size_t readBytes = fread(buffer, 1, (size_t)size, fp);
  if (readBytes != (size_t)size) {
    perror("fread");
    exit(1);
  }
  buffer[size] = '\0';
  fclose(fp);

  return buffer;
}

void part_1(const char *filepath) {
  char *data = read_file(filepath);

  int count = 0;
  int pointer = 50;

  char *token = strtok(data, " \t\n");

  while (token != NULL) {
    char dir = token[0];
    int dist = atoi(token + 1);

    switch (dir) {
    case 'L':
      pointer = (pointer - dist + 100) % 100;
      break;
    case 'R':
      pointer = (pointer + dist) % 100;
      break;
    default:
      fprintf(stderr, "invalid direction: %c\n", dir);
      free(data);
      exit(1);
    }

    if (pointer == 0) {
      count++;
    }

    token = strtok(NULL, " \t\n");
  }

  printf("%d\n", count);

  free(data);
}

void part_2(const char *filepath) {
  char *data = read_file(filepath);

  int count = 0;
  int pointer = 50;

  char *token = strtok(data, " \t\n");
  while (token != NULL) {

    char dir = token[0];
    int dist = atoi(token + 1);

    if (dist == 0 && token[1] != '0') {
      fprintf(stderr, "invalid line: %s\n", token);
      free(data);
      exit(1);
    }

    div_t res = div(dist, 100);
    count += res.quot;

    switch (dir) {
    case 'L':
      if (pointer != 0 && res.rem >= pointer)
        count++;
      pointer = (pointer - res.rem + 100) % 100;
      break;

    case 'R':
      if (pointer + res.rem >= 100)
        count++;
      pointer = (pointer + res.rem) % 100;
      break;

    default:
      fprintf(stderr, "invalid direction: %c\n", dir);
      free(data);
      exit(1);
    }

    token = strtok(NULL, " \t\n");
  }

  printf("%d\n", count);

  free(data);
}
