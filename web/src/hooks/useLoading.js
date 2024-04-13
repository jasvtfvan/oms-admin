import { ref, computed } from 'vue';

const loadingValue = ref(false);
const tipValue = ref('');

export function useLoading() {
  function showLoading(tip) {
    loadingValue.value = true;
    tipValue.value = tip;
  }
  function hideLoading() {
    loadingValue.value = false;
    tipValue.value = '';
  }

  const getLoading = computed(() => loadingValue.value);
  const getLoadingTip = computed(() => tipValue.value);

  return { getLoading, getLoadingTip, showLoading, hideLoading };
}
