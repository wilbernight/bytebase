<template>
  <div class="space-y-4 divide-y divide-block-border">
    <BrandingSetting />
    <SecuritySetting />
    <NetworkSetting v-if="!isSaaSMode" />
    <AIAugmentationSetting />
  </div>
</template>

<script lang="ts" setup>
import { storeToRefs } from "pinia";
import {
  BrandingSetting,
  SecuritySetting,
  NetworkSetting,
  AIAugmentationSetting,
} from "@/components/GeneralSetting";
import { useActuatorStore } from "@/store";

const actuatorStore = useActuatorStore();

const { isSaaSMode } = storeToRefs(actuatorStore);
</script>
