{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "marathon.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "marathon.fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "marathonserver.fullname" -}}
{{- printf "%s-%s" .Release.Name "marathonserver" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "prometheus.fullname" -}}
{{- printf "%s-%s" .Release.Name "marathonserver" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
