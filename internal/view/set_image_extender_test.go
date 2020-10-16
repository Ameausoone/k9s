package view

import "testing"

func Test_imageFormSpec_modified(t *testing.T) {
	var data = imageFormSpec{dockerImage: "foo"}

	if data.modified() != false {
		t.Error("new name empty expected modified false")
	}

	data.newDockerImage = "  "
	if data.modified() != false {
		t.Error("new name composed of spaces expected modified false")
	}

	data.newDockerImage = "bar"
	if data.modified() != false {
		t.Errorf("new name '%s' expected modified false", data.newDockerImage)
	}

	data.newDockerImage = " foo "
	if data.modified() != true {
		t.Errorf("new name '%s' expected modified true", data.newDockerImage)
	}

	data.newDockerImage = "foo"
	if data.modified() != true {
		t.Errorf("new name '%s' expected modified true", data.newDockerImage)
	}
}

func Test_imageFormSpec_imageSpec(t *testing.T) {
	var data = imageFormSpec{dockerImage: "foo"}

	var dockerImageExpected = "foo"
	var initExpected = false
	if spec := data.imageSpec(); spec.DockerImage != dockerImageExpected {
		t.Errorf("docker image expected %s - get %s", dockerImageExpected, spec.DockerImage)
	} else if spec.Init != initExpected {
		t.Errorf("init expected %v - get %v", initExpected, spec.Init)
	}

	data.init = true
	initExpected = true
	if spec := data.imageSpec(); spec.Init != initExpected {
		t.Errorf("init expected %v - get %v", initExpected, spec.Init)
	}

	data.newDockerImage = "  "
	if spec := data.imageSpec(); spec.DockerImage != dockerImageExpected {
		t.Errorf("docker image expected %s - get %s", dockerImageExpected, spec.DockerImage)
	}

	dockerImageExpected = "bar"
	data.newDockerImage = " bar "
	if spec := data.imageSpec(); spec.DockerImage != dockerImageExpected {
		t.Errorf("docker image expected %s - get %s", dockerImageExpected, spec.DockerImage)
	}

	data.newDockerImage = "bar"
	if spec := data.imageSpec(); spec.DockerImage != dockerImageExpected {
		t.Errorf("docker image expected %s - get %s", dockerImageExpected, spec.DockerImage)
	}
}
