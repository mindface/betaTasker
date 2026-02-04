import { Mesh } from "three";

export type VRMSpringBoneColliderMesh = Mesh;

export interface VRMSpringBoneColliderGroup {
  node: number;
  colliders: VRMSpringBoneColliderMesh[];
}

export namespace VRMSchema {
  export interface VRM {
    blendShapeMaster?: BlendShape;
    experterVersion?: string;
    firstPerson?: FirstPerson;
    humanoid?: Humanoid;
    materialProperties?: Material[];
    meta?: Meta;
    secondaryAnimation?: SecondaryAnimation;
    specVersion?: string;
  }

  export interface BlendShape {
    blendShapeGroups?: BlendShapeGroup[];
  }

  export interface BlendShapeGroup {
    binds?: BlendShapeBind[];
    isBinary?: boolean;
    materialValues?: BlendShapeMaterialbind[];
    name?: string;
    presetName?: BlendShapePresetName;
  }

  export interface BlendShapeBind {
    index?: number;
    mesh?: number;
    weight?: number;
  }

  export interface BlendShapeMaterialbind {
    materialName?: string;
    propertyName?: string;
    targetValue?: number[];
  }

  export enum BlendShapePresetName {
    A = "a",
    Angry = "angry",
    Blink = "blink",
    BlinkL = "blink_l",
    BlinkR = "blink_r",
    E = "e",
    Fun = "fun",
    I = "i",
    Joy = "joy",
    LookDown = "lookdown",
    Lookleft = "lookleft",
    Lookright = "lookright",
    Lookup = "lookup",
    Neutral = "neutral",
    O = "o",
    Sorrow = "sorrow",
    U = "u",
    Unknown = "unknown",
  }

  export interface FirstPerson {
    firstPersonBone?: number;
    firstPersonBoneOffset?: Vector3;
    lookAtHorizontalInner?: FirstPersonDegreeMap;
    lookAtHorizontalOuter?: FirstPersonDegreeMap;
    lookAtTypeName?: FirstPersonLookAtTypeName;
    lookAtVerticalDown?: FirstPersonDegreeMap;
    lookAtVerticalUp?: FirstPersonDegreeMap;

    meshAnnotations?: FirstPersonMeshannotation[];
  }

  export interface FirstPersonDegreeMap {
    curve?: number[];
    xRange?: number;
    yRange?: number;
  }

  export enum FirstPersonLookAtTypeName {
    BlendShape = "BlendShape",
    Bone = "Bone",
  }

  export interface FirstPersonMeshannotation {
    firstPersonFlag?: string;
    mesh?: number;
  }

  export interface Humanoid {
    armStretch?: number;
    feetSpacing?: number;
    hasTranslationDof?: boolean;
    humanBones?: HumanoidBone[];
    legStretch?: number;
    lowerArmTwist?: number;
    lowerLegTwist?: number;
    upperArmTwist?: number;
    upperLegTwist?: number;
  }

  export interface HumanoidBone {
    axisLegnth?: number;
    bone?: HumanoidBoneName;
    center?: Vector3;
    max?: Vector3;
    mix?: Vector3;
    node?: number;
    useDefaultValues?: boolean;
  }

  export enum HumanoidBoneName {
    Chest = "chest",
    Head = "head",
    Hips = "hips",
    jaw = "jaw",
    LeftEye = "leftEye",
    LeftFoot = "leftFoot",
    LeftHand = "leftHand",
    LeftIndexDistal = "leftIndexDistal",
    LeftIndexIntermediate = "leftIndexIntermediate",
    LeftIndexProximal = "leftIndexProximal",
    LeftLittleIntermediate = "leftLittleIntermediate",
    LeftLittleProximal = "leftLittleProximal",
    LeftLowerArm = "leftLowerArm",
    LeftLowerLeg = "leftLowerLeg",
    LeftMiddleDistal = "leftMiddleDistal",
    LeftMiddleIntermediate = "leftMiddleIntermediate",
    LeftMiddleProximal = "leftMiddleProximal",
    LeftRingDistal = "leftRingDistal",
    LeftRingIntermediate = "leftRingIntermediate",
    LeftRingProximal = "leftRingProximal",
    LeftShoulder = "leftShoulder",
    LeftThumbDistal = "leftThumbDistal",
    LeftThumbIntermediate = "leftThumbIntermediate",
    LeftThumbProximal = "leftThumbProximal",
    LeftToes = "leftToes",
    LeftUpperArm = "leftUpperArm",
    LeftUpperLeg = "leftUpperLeg",
    Neck = "neck",
    RightEye = "rightEye",
    RightFoot = "rightFoot",
    RightHand = "rightHand",
    RightIndexDistal = "rightIndexDistal",
    RightIndexIntermediate = "rightIndexIntermediate",
    RightIndexProximal = "rightIndexProximal",
    RightLittleDistal = "rightLittleDistal",
    RightLittleIntermediate = "rightLittleIntermediate",
    RightLittleProximal = "rightLittleProximal",
    RightLowerArm = "rightLowerArm",
    RightLowerLeg = "rightLowerLeg",
    RightMiddleDistal = "rightMiddleDistal",
    RightMiddleIntermediate = "rightMiddleIntermediate",
    RightMiddleProximal = "rightMiddleProximal",
    RightRingDistal = "rightRingDistal",
    RightRingIntermediate = "rightRingIntermediate",
    RightRingProximal = "rightRingProximal",
    RightShoulder = "rightShoulder",
    RightThumbDistal = "rightThumbDistal",
    RightThumbIntermediate = "rightThumbIntermediate",
    RightThumbProximal = "rightThumbProximal",
    RightToes = "rightToes",
    RightUpperArm = "rightUpperArm",
    RightUpperLeg = "rightUpperLeg",
    Spine = "spine",
    UpperChest = "upperChest",
  }

  export interface Material {
    floatProperties?: { [key: string]: any };
    keywordMap?: { [key: string]: any };
    name?: string;
    renderQueue?: number;
    shander?: string;
    tagMap?: { [key: string]: any };
    textureProperties?: { [key: string]: any };
    vectorProperties?: { [key: string]: any };
  }

  export interface Meta {
    allowedUserName?: MetaAllowedUserName;
    author?: string;
    commercialUssageName?: MetaUssageName;
    contactInformation?: string;
    licenseName?: MetaLicenseName;
    otherLicenseUrl?: string;
    otherPermissionUrl?: string;
    reference?: string;
    sexualUssageName?: MetaUssageName;
    texture?: number;
    title?: string;
    version?: string;
    violentUssageName?: MetaUssageName;
  }

  export enum MetaAllowedUserName {
    Everyone = "Everyone",
    ExplicitlyLicensedPerson = "ExplicitlyLicensedPerson",
    OnlyAuthor = "OnlyAuthor",
  }

  export enum MetaUssageName {
    Allow = "Allow",
    Disallow = "Disallow",
  }

  export enum MetaLicenseName {
    Cc0 = "Cc0",
    CcBy = "CC_BY",
    CcByNc = "CC_BY_NC",
    CcByNcNd = "CC_BY_NC_ND",
    CcByNcSa = "CC_BY_NC_SA",
    CcByNd = "CC_BY_ND",
    CcBySa = "CC_BY_SA",
    Other = "other",
    RedistributionProhibited = "Redistribution_Prohibited",
  }

  export interface SecondaryAnimation {
    boneGroups?: SecondaryAnimationSpring[];
    colliderGroups?: SecondaryAnimationCollidergroup[];
  }

  export interface SecondaryAnimationSpring {
    bones?: number[];
    center?: number;
    colliderGroups?: number[];
    comment?: string;
    dragForce?: number;
    gravityDir?: Vector3;
    gravityPower?: number;
    hitRadius?: number;
    stiffiness?: number;
  }

  export interface SecondaryAnimationCollidergroup {
    colliders?: SecondaryAnimationCollider[];
    node?: number;
  }

  export interface SecondaryAnimationCollider {
    offset?: Vector3;
    radius?: number;
  }

  export interface Vector3 {
    x?: number;
    y?: number;
    z?: number;
  }
}
