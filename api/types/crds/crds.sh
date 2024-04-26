CRDS=$(ls -1 *crd.yaml | grep -v skupper_cluster_policy)
for crd in ${CRDS[@]}; do
    name=$(cat ${crd} | yq -r .metadata.name)
    kind=$(cat ${crd} | yq -r .spec.names.kind)
    desc="Skupper ${kind}"
    #kind=$(echo "${crd}" | awk -F. '{print $1}' | sed -re 's/(.*)s$/\1/g')
    #[[ "${kind}" =~ ccesse$ ]] && kind=$(echo "${crd}" | awk -F. '{print $1}' | sed -re 's/(.*)ccesses$/\1ccess/g')
    #display_name=$(echo "${kind}" | awk '{print toupper(substr( $0, 1, 1 ) ) substr( $0, 2 );}')
    cat << EOF
    - name: ${name}
      group: skupper.io
      description: ${desc}
      displayName: ${kind}
      kind: ${kind}
      version: v1alpha1
EOF
done
