// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: efg/backstage/service_options.proto

package com.efg.backstage;

public final class ServiceOptions {
  private ServiceOptions() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
    registry.add(com.efg.backstage.ServiceOptions.owner);
    registry.add(com.efg.backstage.ServiceOptions.system);
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public static final int OWNER_FIELD_NUMBER = 51001;
  /**
   * <code>extend .google.protobuf.ServiceOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.ServiceOptions,
      java.lang.String> owner = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        java.lang.String.class,
        null);
  public static final int SYSTEM_FIELD_NUMBER = 51002;
  /**
   * <code>extend .google.protobuf.ServiceOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.ServiceOptions,
      java.lang.String> system = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        java.lang.String.class,
        null);

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n#efg/backstage/service_options.proto\022\re" +
      "fg.backstage\032 google/protobuf/descriptor" +
      ".proto::\n\005owner\022\037.google.protobuf.Servic" +
      "eOptions\030\271\216\003 \001(\tR\005owner\210\001\001:<\n\006system\022\037.g" +
      "oogle.protobuf.ServiceOptions\030\272\216\003 \001(\tR\006s" +
      "ystem\210\001\001BV\n\021com.efg.backstageZAgithub.co" +
      "m/faceit/protoc-gen-backstage/proto/gen/" +
      "go/efg/backstageb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.DescriptorProtos.getDescriptor(),
        });
    owner.internalInit(descriptor.getExtensions().get(0));
    system.internalInit(descriptor.getExtensions().get(1));
    com.google.protobuf.DescriptorProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
