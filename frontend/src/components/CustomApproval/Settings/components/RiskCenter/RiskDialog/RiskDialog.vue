<template>
  <BBModal
    v-if="dialog"
    :title="title"
    :esc-closable="false"
    :before-close="beforeClose"
    :data-state-dirty="state.dirty"
    @close="dialog = undefined"
  >
    <RiskForm
      :dirty="state.dirty"
      @cancel="cancel"
      @update="state.dirty = true"
      @save="handleSave"
    />
    <div
      v-if="state.loading"
      class="absolute inset-0 flex flex-col items-center justify-center bg-white/50 rounded-lg"
    >
      <BBSpin />
    </div>
  </BBModal>
</template>

<script lang="ts" setup>
import { computed, reactive, watch } from "vue";
import { useI18n } from "vue-i18n";
import { useDialog } from "naive-ui";

import { defer } from "@/utils";
import { useRiskCenterContext } from "../context";
import { sourceText } from "../../common";
import RiskForm from "./RiskForm.vue";
import { Risk } from "@/types/proto/v1/risk_service";
import { pushNotification, useRiskStore } from "@/store";

type LocalState = {
  loading: boolean;
  dirty: boolean;
};

const { t } = useI18n();
const context = useRiskCenterContext();
const { allowAdmin, dialog } = context;
const nDialog = useDialog();

const state = reactive<LocalState>({
  loading: false,
  dirty: false,
});

const title = computed(() => {
  const parts: string[] = [];
  if (dialog.value) {
    const { mode, risk } = dialog.value;
    if (!allowAdmin.value) {
      parts.push(t("custom-approval.security-rule.view-rule"));
    } else {
      if (mode === "CREATE") {
        parts.push(t("custom-approval.security-rule.add-rule"));
      } else if (mode === "EDIT") {
        parts.push(t("custom-approval.security-rule.edit-rule"));
      }
    }
    parts.push(sourceText(risk.source));
  }
  return parts.join(" - ");
});
const cancel = async () => {
  const pass = await beforeClose();
  if (pass) {
    dialog.value = undefined;
  }
};

const beforeClose = async () => {
  if (!state.dirty) {
    return true;
  }
  if (!allowAdmin.value) {
    return true;
  }
  const d = defer<boolean>();
  nDialog.info({
    title: t("common.close"),
    content: t("common.will-lose-unsaved-data"),
    maskClosable: false,
    closeOnEsc: false,
    positiveText: t("common.confirm"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => d.resolve(true),
    onNegativeClick: () => d.resolve(false),
  });
  return d.promise;
};

const handleSave = async (risk: Risk) => {
  state.loading = true;

  try {
    await useRiskStore().upsertRisk(risk);

    state.dirty = false;
    pushNotification({
      module: "bytebase",
      style: "SUCCESS",
      title: t("common.updated"),
    });
    dialog.value = undefined;
  } finally {
    state.loading = false;
  }
};

watch(
  dialog,
  () => {
    state.dirty = false;
  },
  { immediate: true }
);
</script>
