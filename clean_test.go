package clean_sec_group

import "testing"

func TestRemoveEnd(t *testing.T) {
	var edge [2]string
	edge[0] = "A"
	edge[1] = "B"
	if edge[0] == "A" {
		t.Log("edge[0]==A")
	}
	if edge[1] == "B" {
		t.Log("edge[1]==B")
	}
	fixed, err := remove_sec(edge, "B")
	if fixed == true && err == nil {
		t.Log("sec removed")
	} else {
		t.Logf("sec removed %t %s", fixed, err)
	}
}
func TestRemoveNotEnd(t *testing.T) {
	var edge [2]string
	edge[0] = "A"
	edge[1] = "B"
	if edge[0] == "A" {
		t.Log("edge[0]==A")
	}
	if edge[1] == "B" {
		t.Log("edge[1]==B")
	}
	fixed, err := remove_sec(edge, "A")
	if fixed == true && err == nil {
		t.Log("sec removed")
	} else {
		t.Errorf("sec was not be removed %s", err)
	}
}

// aws ec2 describe-security-groups --query "SecurityGroups[*].{ID:GroupId,Name:GroupName,dependentOnSGs:IpPermissions[].UserIdGroupPairs[].GroupId}"
