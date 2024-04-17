import { union, isEqual, unionWith } from 'lodash-es';

export const mergeArray = (arr1, arr2) => {
  return union(arr1, arr2)
}

export const mergeArrayDeep = (arr1, arr2) => {
  return unionWith(arr1, arr2, isEqual)
}
